import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as lambda from 'aws-cdk-lib/aws-lambda'
import * as apigw from 'aws-cdk-lib/aws-apigateway'

export class GenUrlIdStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const genUrlIdLambda = new lambda.Function(this, 'GenUrlIdLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2023,
      code: lambda.Code.fromAsset('../deployment/deployment.zip'),
      handler: 'bootstrap',
      architecture: lambda.Architecture.ARM_64,
    });

    const api = new apigw.LambdaRestApi(this, 'GenUrlApi', {
      handler: genUrlIdLambda,
    });
  }
}
