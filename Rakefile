require 'tempfile'
require 'yaml'
require 'colorize'

namespace :pipeline do
  desc 'Login to concourse'
  task :setup do
  end

  desc 'Login to concourse'
  task :login do
    exec('fly -t osl login -c https://osl.ci.cf-app.com/ osl')
  end

  desc 'Trigger build on concourse'
  task :trigger, [:config_file, :environment] do |_, args|
    config_file = args[:config_file]
    suffix = args[:environment] == 'staging' ? '-staging': ''
    pipeline_name = File.basename(config_file, '.yml') + suffix

    sh("fly -t osl trigger-job -j #{pipeline_name}/dependency-report")
  end

  desc 'Configure a single pipeline'
  task :configure, [:config_file, :environment] do |_, args|
    config_file = args[:config_file]

    configure_report_pipelines [config_file], args[:environment]
  end

  desc 'Configure all concourse pipelines'
  task :configure_all, [:environment] do |_, args|
    config_files = Dir[File.join('releases', '*.yml')]
    configure_report_pipelines config_files, args[:environment]
  end

  desc 'Delete all staging pipelines'
  task :delete_all_staging do
    pipelines = `fly -t osl pipelines | sed -E 's/[ ]+(yes|no)[ ]+(yes|no)$//' | grep staging`.split("\n")

    pipelines.each do |pipeline|
      cmd = "fly -t osl dp -n -p #{pipeline}"
      puts cmd
      IO.popen(cmd).each do |line|
        puts line.chomp
      end
    end
  end
end

def configure_report_pipelines(config_files, environment)
  config_files.each do |config_file|
    set_pipeline(config_file, environment)
  end
end

def check_lastpass
  begin
    sh('lpass ls > /dev/null')
  rescue
    puts 'Login required with lpass'
    exit 1
  end
end

def set_pipeline(config_file, environment)
  suffix = environment == 'staging' ? '-staging' : ''

  product_name = File.basename(config_file, '.yml')
  pipeline_name = "#{product_name}#{suffix}"
  puts "Updating pipeline #{pipeline_name}".blue

  if environment == 'staging'
    osl_branch = 'master'
  elsif environment == 'production'
    osl_branch = 'passed'
  else
    raise "unknown environment: #{environment} (try staging or production)"
  end

  vars = ["release-config=#{config_file}"]
  vars.push "osl-branch=#{osl_branch}"

  release_file = Tempfile.new('release.yml')
  begin
    release_file.write generate_pipeline(config_file)
    release_file.close
    command = fly_command(environment, pipeline_name, release_file, vars)
    sh(command, verbose: true)
    raise 'updating the pipeline failed.' unless $?.success?
  ensure
    release_file.unlink
  end
end


def fly_command(environment, pipeline_name, release_file, vars)
  check_lastpass
  var_params = vars.map do |var|
    "--var #{var}"
  end.join(' ')

  "bash -c \"fly -t osl set-pipeline -p #{pipeline_name} -c #{release_file.path} \
     --load-vars-from <(lpass show 'Shared-Open Source Licensing Tooling /osl-ci-private-fly' --notes) \
     --load-vars-from <(lpass show 'Shared-Open Source Licensing Tooling /osl-ci-private-#{environment}' --notes) \
     #{var_params} -n\""
end

def generate_pipeline(config_file)
  config_yml = YAML.load File.read(config_file)

  product_name = File.basename(config_file, '.yml')
  releases = config_yml['releases']

  releases.each do |release|
    if !release['name'] && release['repo']
      release['name'] = repo_name_from_url(release['repo'])
    elsif !release['name'] && release['stemcell_name']
      release['name'] = release['stemcell_name']
    end
  end

  resource_types = [
      create_slack_resource_type
  ]

  resources = [
      create_slack_resource,
      create_osl_ci_resource,
      create_osl_config_resource
  ]
  jobs = []

  releases.each do |release|
    unless release['repo'].to_s.empty?
      resources << create_git_resource(release)
    end

    unless release['stemcell_name'].to_s.empty?
      resources << create_stemcell_resource(release)
    end

    resources << create_dependency_report_s3_resource(product_name, release)

    resources << create_release_report_s3_resource(product_name, release)

    jobs << create_dependency_report_job(release)

    jobs << create_release_report_job(release)
  end

  pipeline = {
      'resource_types' => resource_types,
      'resources' => resources,
      'jobs' => jobs
  }

  pipeline_yaml = pipeline.to_yaml

  # ruby yaml does not like parenthesis without wrapping it in quotes,
  # but we need it in order to work with fly string replacement
  pipeline_yaml.gsub('"{{', '{{').gsub('}}"', '}}')
end


def create_slack_resource_type
  {
      'name' => 'slack-notification',
      'type' => 'docker-image',
      'source' => {
          'repository' => 'cfcommunity/slack-notification-resource',
          'tag' => 'latest'
      }
  }
end

def create_git_resource(release)
  {
      'name' => "#{release['name']}-git",
      'type' => 'git',
      'source' => {
          'uri' => release['repo'],
          'private_key' => '{{CfOslBotPrivateKey}}',
          'branch' => release['branch']
      }
  }
end

def create_stemcell_resource(release)
  {
      'name' => "#{release['name']}-stemcell",
      'type' => 'bosh-io-stemcell',
      'source' => {
          'name' => release['stemcell_name']
      }
  }
end

def create_dependency_report_s3_resource(product_name, release)
  {
      'name' => "#{release['name']}-s3-dependency-reports",
      'type' => 's3',
      'source' => {
          'bucket' => '{{S3ReportBucket}}',
          'access_key_id' => '{{AwsAccessKeyId}}',
          'secret_access_key' => '{{AwsSecretAccessKey}}',
          'private' => true,
          'versioned_file' => "#{product_name}/#{release['name']}/dependency_reports/reports.tar.gz"
      }
  }
end

def create_release_report_s3_resource(product_name, release)
  {
      'name' => "#{release['name']}-s3-release-reports",
      'type' => 's3',
      'source' => {
          'bucket' => '{{S3ReportBucket}}',
          'access_key_id' => '{{AwsAccessKeyId}}',
          'secret_access_key' => '{{AwsSecretAccessKey}}',
          'private' => true,
          'versioned_file' => "#{product_name}/#{release['name']}/release-report.xml"
      }
  }
end

def create_dependency_report_job(release)
  release_resource = "#{release['name']}-git"
  if release['stemcell_name']
    release_resource = "#{release['name']}-stemcell"
  end

  {
      'name' => "#{release['name']}-dependency-report",
      'plan' => [
          {
              'aggregate' => [
                  {'get' => 'osl-ci'},
                  {'get' => 'osl-config'},
                  {
                      'get' => 'release',
                      'resource' => release_resource,
                      'trigger' => true
                  }
              ]
          },
          {
              'task' => 'report',
              'file' => 'osl-ci/ci/tasks/dependency_report/task.yml',
              'params' => {
                  'RELEASE_NAME' => "#{release['name']}",
                  'ENDPOINT' => '{{endpoint}}',
                  'CONFIG_FILE' => '{{release-config}}',
                  'GITHUB_KEY' => '{{CfOslBotPrivateKey}}',
                  'GITHUB_API_USER' => '{{GithubApiUser}}',
                  'GITHUB_API_TOKEN' => '{{GithubApiTokenPipelines}}',
                  'GITHUB_API_ROOT_URL' => 'https://api.github.com'
              }
          },
          {
              'put' => "#{release['name']}-s3-dependency-reports",
              'params' => {
                  'file' => 'reports/reports.tar.gz'
              }
          }
      ],
      'on_failure' => {
          'put' => 'slack-alert',
          'params' => {
              'channel' => '{{channel}}',
              'icon_emoji' => ':thumbsdown:',
              'text' => "Dependency report generation failed for #{release['name']}\nhttps://osl.ci.cf-app.com/teams/main/pipelines/$BUILD_PIPELINE_NAME/jobs/$BUILD_JOB_NAME/builds/$BUILD_NAME"
          }
      }
  }
end

def create_release_report_job(release)
  release_resource = "#{release['name']}-git"
  if release['stemcell_name']
    release_resource = "#{release['name']}-stemcell"
  end
  {
      'name' => "#{release['name']}-release-report",
      'plan' => [
          {
              'aggregate' => [
                  {'get' => 'osl-ci'},
                  {'get' => 'osl-config'},
                  {
                      'get' => 'release',
                      'resource' => release_resource,
                      'passed' => ["#{release['name']}-dependency-report"],
                      'trigger' => true
                  }
              ]
          },
          {
              'task' => 'report',
              'file' => 'osl-ci/ci/tasks/release_report/task.yml',
              'params' => {
                  'RELEASE_NAME' => "#{release['name']}",
                  'ENDPOINT' => '{{endpoint}}',
                  'CONFIG_FILE' => '{{release-config}}',
                  'GITHUB_KEY' => '{{CfOslBotPrivateKey}}',
              }
          },
          {
              'put' => "#{release['name']}-s3-release-reports",
              'params' => {
                  'file' => 'report/release-report.xml'
              }
          }
      ],
      'on_failure' => {
          'put' => 'slack-alert',
          'params' => {
              'channel' => '{{channel}}',
              'icon_emoji' => ':thumbsdown:',
              'text' => "Release report generation failed for #{release['name']}\nhttps://osl.ci.cf-app.com/teams/main/pipelines/$BUILD_PIPELINE_NAME/jobs/$BUILD_JOB_NAME/builds/$BUILD_NAME"
          }
      }
  }
end

def create_osl_ci_resource
  {
      'name' => 'osl-ci',
      'type' => 'git',
      'source' => {
          'uri' => 'git@github.com:pivotal-cf/osl-ci.git',
          'private_key' => '{{CfOslBotPrivateKey}}',
          'branch' => '{{osl-branch}}'
      }
  }
end

def create_osl_config_resource
  {
      'name' => 'osl-config',
      'type' => 'git',
      'source' => {
          'uri' => 'git@github.com:pivotal-cf/osl-config.git',
          'private_key' => '{{CfOslBotPrivateKey}}'
      }
  }
end

def create_slack_resource
  {
      'name' => 'slack-alert',
      'type' => 'slack-notification',
      'source' => {
          'url' => 'https://hooks.slack.com/services/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'
      }
  }
end

def sh(cmd, opts = {})
  puts "Running #{cmd}".blue if opts[:verbose]
  output = []
  IO.popen(cmd).each do |line|
    puts line.chomp if opts[:verbose]
    output << line.chomp
  end
  raise "#{cmd} failed: #{output}" unless $?.success?
  output.join("\n")
end

def repo_name_from_url(repo_url)
  remote = repo_url.chomp('.git')
  File.basename remote
end
