package main

import (
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"log"
	"os"
	"strings"
	"time"
)

func addMetric(name, unit string, value float64, dimensions []*cloudwatch.Dimension, metricData []*cloudwatch.MetricDatum) (ret []*cloudwatch.MetricDatum, err error) {
	_metric := cloudwatch.MetricDatum{
		MetricName: aws.String(name),
		Unit:       aws.String(unit),
		Value:      aws.Float64(value),
		Dimensions: dimensions,
		Timestamp:  aws.Time(time.Now()),
	}
	metricData = append(metricData, &_metric)
	return metricData, nil
}

func putMetric(metricdata []*cloudwatch.MetricDatum, namespace, region string) error {
	svc := cloudwatch.New(session.New(), &aws.Config{Region: aws.String(region)})

	metric_input := &cloudwatch.PutMetricDataInput{
		MetricData: metricdata,
		Namespace:  aws.String(namespace),
	}

	resp, err := svc.PutMetricData(metric_input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			log.Println("Error when putMetric, ", awsErr.Code(), awsErr.Message())
			log.Println("Error:", awsErr.Error())
			return err
		} else if err != nil {
			return err
		}
	}
	log.Println(awsutil.StringValue(resp))
	return nil
}

func main() {
	var metricData []*cloudwatch.MetricDatum
	var dims []*cloudwatch.Dimension
	metric_name := flag.String("metric-name", "", "Cloudwatch metric name(required)")
	ns := flag.String("namespace", "Linux/System", "Cloudwatch metric namespace(default: Linux/System)")
	unit := flag.String("unit", "Count", "Metric Unit")
	ds := flag.String("dimensions", "", "Cloudwatch Dimensions")
	region := flag.String("region", "us-east-1", "AWS Region")
	val := flag.Float64("value", 0.0, "Value of the metric")

	flag.Parse()

	if *metric_name == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	for _, d := range strings.Split(*ds, ",") {
		_ds := strings.Split(d, "=")
		dim := cloudwatch.Dimension{
			Name:  aws.String(_ds[0]),
			Value: aws.String(_ds[1]),
		}
		dims = append(dims, &dim)
	}
	var err error
	metricData, err = addMetric(*metric_name, *unit, *val, dims, metricData)
	if err != nil {
		log.Fatal("Can't add metric")
	}
	err = putMetric(metricData, *ns, *region)
	if err != nil {
		log.Fatal("Can't put CloudWatch Metric")
	}
}
