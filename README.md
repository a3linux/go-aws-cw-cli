Simple AWS CW Cli binary command
=================================

* To put metric
```
go-aws-cw-cli --metric-name <metric-name> --namespace <namespace> --dimensions "Dimension1=Value,Dimension2=Value2" --value <val> [ --unit Count --region <region> ]

Usage
  -dimensions string
  Cloudwatch Dimensions
  -metric-name string
  Cloudwatch metric name(required)
  -namespace string
  Cloudwatch metric namespace(default: Linux/System) (default "Linux/System")
  -region string
  AWS Region (default "us-east-1")
  -unit string
  Metric Unit (default "Count")
  -value float
  Value of the metric
```

