## About this project

OSLV is a project to detect Open Source License usage in products.
This repo contains the config files for those products.

## Concourse

- Concourse is located at [osl.ci.cf-app.com](https://osl.ci.cf-app.com)
- Authentication uses github and all `Open Source License Tooling`
  members should have access
  
### Updating the pipeline

1. Install the `lpass` cli if you haven't already
1. Make sure to login to run `lpass login USERNAME`
1. From `osl-ci` run the following:
	1. Run `rake pipeline:login` to login to concourse
	1. Run `rake pipeline:configure_all` to configure all pipelines
        1. Optionally, run `rake pipeline:configure[config_file,environment]` to configure a single pipeline

eg.

    rake pipeline:configure[releases/apps-manager-release.yml,production]

or

    rake pipeline:configure[releases/apps-manager-release.yml,staging]

In order to add/remove a project please edit `releases/*.yml`. Here's an example from `diego-windows-release.yml`:

```
---
repo: https://github.com/cloudfoundry/diego-windows-release
branch: master
projects:
- name: metron
  path: loggregator/src/metron
- name: DiegoWindowsRelease
  path: DiegoWindowsRelease
```

The submodules of the the repo are automatically scanned so they don't need to be specified manually. The `projects`
array is only necessary for projects that are not submodules. After creating a new configuration, use
`rake pipeline:configure` and point to the configuration.

#### Prepare the Path for `license_finder`

Use the optional `prepare_command` field of `projects` to specify a command to run before the `license_finder` is
executed. This custom prepare command will run inside the directory specified by `path` and is exectued right before
`license_finder` is run.

Here's an example from `push-analytics.yml`. The prepare command will `pushd` into the parent directory where
a maven `pom.xml` exists that needs to be installed. This is necessary for `license_finder` to run properly later.

```
---
repo: git@github.com:cfmobile/push-analytics.git
branch: dev
scan_root: false
projects:
- name: push-analytics
  path: push-analytics
  prepare_command: bash -c "pushd .. && mvn clean install -DskipTests && popd"
```

### Trigger all job for a pipeline

There is an open [PR](https://github.com/concourse/fly/pull/58) to add support
for triggering all jobs within a pipeline. When merged you run the following to
trigger all jobs under a pipeline.

    rake pipeline:trigger[config_file, environment]

eg.

    rake pipeline:trigger[releases/apps-manager-release.yml,staging]

### Hijacking a build

    fly -t osl hijack -j cf-release/dependency-report -s report

