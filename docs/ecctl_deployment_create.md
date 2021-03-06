## ecctl deployment create

Creates a deployment

### Synopsis

Creates a deployment which can be defined through flags or from a file definition.
Sane default values are provided, making the command work out of the box even when no parameters are set. 
When version is not specified, the latest available stack version will automatically be used. 
These are the available options:

  * Simplified flags to set size and zone count for each instance type (Elasticsearch, Kibana, APM and AppSearch). 
  * Advanced flag for different Elasticsearch node types: --topology-element <json obj> (shorthand: -e).
    Note that the flag can be specified multiple times for complex topologies.
    The JSON object has the following format:
    {
      "name": "["data", "master", "ml"]" # type string
      "size": 1024 # type int32
      "zone_count": 1 # type int32
    }
  * File definition: --file=<file path> (shorthand: -f). You can create a definition by using the sample JSON seen here:
    https://elastic.co/guide/en/cloud/current/ec-api-deployment-crud.html#ec_create_a_deployment

As an option "--generate-payload" can be used in order to obtain the generated payload that would be sent as a request. 
Save it, update or extend the topology and create a deployment using the saved payload with the "--file" flag.

```
ecctl deployment create {--file | --size <int> --zones <string> | --topology-element <obj>} [flags]
```

### Examples

```
## Create a deployment with the default values for Elasticsearch, a Kibana instance with a modified size, 
and a default APM instance. While Elasticsearch and Kibana come enabled by default, both APM and AppSearch need to be 
enabled by using the "--apm" and "--appsearch" flags. The command will exit after the API response has been returned, without 
waiting until the deployment resources have been created. 
$ ecctl deployment create --name my-deployment --zones 2 --kibana-size 2048 --apm --apm-size 1024

## To make the command wait until the resources have been created use the "--track" flag, which will output 
the current stage on which the deployment resources are in.
$ deployment create --name my-deployment --track
[...]
Cluster [38e0ff5b58a9483c96a98c923b22194e][Elasticsearch]: finished running all the plan steps (Total plan duration: 1m0.911628175s)
Cluster [51178ffc645d48b7859dbf17388a6c35][Kibana]: finished running all the plan steps (Total plan duration: 1m11.246662764s)

## Additionally, a more advanced topology for Elasticsearch can be created through "--topology-element" or shorthand "-e".
The following command will create a deployment with two 1GB Elasticsearch instances of the type "data" and 
one 1GB Elasticsearch instance of the type "ml".
$ ecctl deployment create --name my-deployment --topology-element '{"size": 1024, "zone_count": 2, "name": "data"}' --topology-element '{"size": 1024, "zone_count": 1, "name": "ml"}'

## In order to use the "--deployment-template" flag, you'll need to know which deployment templates ara available to you.
Visit https://elastic.co/guide/en/cloud/current/ec-regions-templates-instances.html.
If you are an Elastic Cloud Enterprise customer, you'll need to run the following command to view your deployment templates:
$ ecctl platform deployment-template list

## Use the "--generate-payload" flag to save the definition to a file for later use.
$ ecctl deployment create --name my-deployment --size 1024 --track --generate-payload --zones 2 > elasticsearch_create_example.json

## Create a deployment through the file definition.
$ ecctl deployment create --file elasticsearch_create_example.json --track

## To retry a when the previous deployment creation failed:
$ ecctl deployment create
The deployment creation returned with an error, please use the displayed idempotency token
to recreate the deployment resources
Idempotency token: GMZPMRrcMYqHdmxjIQkHbdjnhPIeBElcwrHwzVlhGUSMXrEIzVXoBykSVRsKncNb
unknown error (status 500)
$ ecctl deployment create --request-id=GMZPMRrcMYqHdmxjIQkHbdjnhPIeBElcwrHwzVlhGUSMXrEIzVXoBykSVRsKncNb
```

### Options

```
      --apm                            Enables APM for the deployment
      --apm-ref-id string              Optional RefId for the APM deployment (default "main-apm")
      --apm-size int32                 Memory (RAM) in MB that each of the APM instances will have (default 512)
      --apm-zones int32                Number of zones the APM instances will span (default 1)
      --appsearch                      Enables AppSearch for the deployment
      --appsearch-ref-id string        Optional RefId for the AppSearch deployment (default "main-appsearch")
      --appsearch-size int32           Memory (RAM) in MB that each of the AppSearch instances will have (default 2048)
      --appsearch-zones int32          Number of zones the AppSearch instances will span (default 1)
      --deployment-template string     Deployment template ID on which to base the deployment from (default "default")
  -f, --file string                    DeploymentCreateRequest file definition. See help for more information
      --generate-payload               Returns the deployment payload without actually creating the deployment resources
  -h, --help                           help for create
      --kibana-ref-id string           Optional RefId for the Kibana deployment (default "main-kibana")
      --kibana-size int32              Memory (RAM) in MB that each of the Kibana instances will have (default 1024)
      --kibana-zones int32             Number of zones the Kibana instances will span (default 1)
      --name string                    Optional name for the deployment
      --plugin strings                 Additional plugins to add to the Elasticsearch deployment
      --ref-id string                  Optional RefId for the Elasticsearch deployment (default "main-elasticsearch")
      --request-id string              Optional idempotency token - Can be found in the Stderr device when a previous deployment creation failed, for more information see the examples in the help command page
      --size int32                     Memory (RAM) in MB that each of the Elasticsearch instances will have (default 4096)
  -e, --topology-element stringArray   Optional Elasticsearch topology element definition. See help for more information
  -t, --track                          Tracks the progress of the performed task
      --version string                 Version to use, if not specified, the latest available stack version will be used
      --zones int32                    Number of zones the Elasticsearch instances will span (default 1)
```

### Options inherited from parent commands

```
      --apikey string      API key to use to authenticate (If empty will look for EC_APIKEY environment variable)
      --config string      Config name, used to have multiple configs in $HOME/.ecctl/<env> (default "config")
      --force              Do not ask for confirmation
      --format string      Formats the output using a Go template
      --host string        Base URL to use
      --insecure           Skips all TLS validation
      --message string     A message to set on cluster operation
      --output string      Output format [text|json] (default "text")
      --pass string        Password to use to authenticate (If empty will look for EC_PASS environment variable)
      --pprof              Enables pprofing and saves the profile to pprof-20060102150405
  -q, --quiet              Suppresses the configuration file used for the run, if any
      --region string      Elasticsearch Service region
      --timeout duration   Timeout to use on all HTTP calls (default 30s)
      --trace              Enables tracing saves the trace to trace-20060102150405
      --user string        Username to use to authenticate (If empty will look for EC_USER environment variable)
      --verbose            Enable verbose mode
```

### SEE ALSO

* [ecctl deployment](ecctl_deployment.md)	 - Manages deployments

