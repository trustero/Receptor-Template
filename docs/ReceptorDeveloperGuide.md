# Trustero Receptor Developer Guide

## Resources

- Trustero SDK reference: https://pkg.go.dev/github.com/trustero/api/go/receptor_sdk
- Example Receptor code (GitLab): https://github.com/trustero/api/tree/main/go/examples/gitlab_receptor
- Receptor Template: https://github.com/trustero/Receptor-Template

## What Is a Receptor?

A Receptor is code that connects to a 3rd party service provider (such as GitHub), gathers information about how you are using that service, and generates evidence to prove proper precautions are in place for a security audit. The evidence is posted to relevant controls in a Trustero account.

The point is to help someone automatically collect information in a somewhat normalized format. The information could be generic, such as what the password policy is. Or it could be granular, such as which users have MFA enabled. In Trustero’s product, this information is then used to make sure procedures are being followed, such as requiring complex passwords or requiring that everyone use MFA.

Although most Receptors connect to 3rd party service platforms (such as GitHub), they can also be written to collect information programmatically from local machines. This guide will focus on writing a Receptor that connects to a 3rd party service provider.

### Evidence Data Model

Receptors exist to gather information. The information has a hierarchical structure with four levels.

At the top of this hierarchy is a `Service Provider`. This refers to a platform from which we can draw data using a single set of credentials. A good example of this is AWS, which allows one to access multiple service APIs such as EC2 and S3 using the same credentials. Another example is GitHub, which is a platform which provides access to a single service API.

Within a Service Provider is one or more `Services`. For AWS, these services include S3, ECS, RDS, and many more.

Within a Service there are one or more `Service Entities`, which are types of objects, such as Users or Repositories.

The specific objects within a Service Entity are called `Service Entity Instances`, such as a single user or a list of S3 buckets.

EX:

- AWS-Receptor (Service Provider)
  - S3 (Service)
    - Buckets (Service Entity)
      - Logo-storage-bucket (Service Entity Instance)
      - Static-assets-bucket (Service Entity Instance)
  - RDS (Service)
    - Instances (Service Entity)
      - ML-testing-DB (Service Entity Instance)
    - Clusters (Service Entity)
      - production-db-stack (Service Entity Instance)
      - staging-db-stack (Service Entity Instance)

### Gathered Information

Receptors will gather and return this information from the Service Provider:

- Which services are being used?
- Within each entity, what are the instances and their properties?
- What queries were made to the Service Provider to get this information?
- What does this data look like when formatted appropriately?
- What does this data look like as a raw API response?

### Evidence

After a Receptor queries a `Service Provider` for any relevant information, the Receptor will need to send this information back to Trustero in the form of `Evidence`.

Evidence consists of:

- A caption - Used to identify the evidence in the UI
  - Ex: “AWS S3 Bucket Encryption Status”
- A description - To explain the significance of the evidence
  - Ex: “A list of AWS S3 buckets and their encryption settings”
- A table/list of `Service Entity Instances` - To show proof of compliance
  - Logo-storage-bucket : AES-256
- An Entity Type -The type of Service Entity
  - Ex: “Bucket”
- A list of API Sources - API calls made and returned JSON responses
  - Ex:
    - call: “s3.getAllBuckets({accountID: 123445667})”
    - response : “{buckets: [Logo-storage-bucket, …]}”

## Expectations

As a contractor, you will be given these 3 items to help you write your Receptor.

- A template repo with a starting directory structure
  - This is where you will write your Receptor based on the requirements provided by Trustero
- A list of services your Receptor should query within a Service Provider
  - You will be provided services you need to gather information for.
  - Ex:
    - Gather all AWS S3 bucket information, as well as encryption settings for each bucket.
    - Gather all AWS EC2 instances
    - Gather all AWS user along with their MFA and admin status
- A list of evidence that the receptor should produce
  - You will be provided a list of evidence that can be generated via the services you have gathered.
  - Ex:
    - Generate evidence that shows all AWS users and their admin status. The evidence should have the caption “User List”
    - Generate evidence that shows all S3 buckets and their encryption settings. The evidence should have the caption “S3 Asset Inventory”
  - **_IMPORTANT - You will need to keep track of all API calls made to gather evidence. Each evidence object should also include the raw API call made, as well as the raw JSON response from the api call_**

## Writing a Receptor

To get you started, you’ll need to research how to connect to your Receptor. You should note this process down as

### Service Provider Credentials

One of the first things you should do when writing a Receptor is research what credentials and permissions will be needed for the Receptor to connect to the Service Provider.

You will need to create a markdown file with instructions to get the required credentials or grant any permissions needed to connect the Receptor to the Service Provider.

If any permissions need to be granted, they should generally be “read-only”.

Here is an example of what the instructions for AWS would look like:

```
In order to use AWS, you must provide AWS API credentials. We recommend
you create a read-only user for this purpose. Follow these steps to create an IAM user:

Log into the AWS IAM console
Select Users from the vertical menu column on the left side of the screen
Select Add User
Enter trustero-api-user in the User name field
Select Programmatic access as the Access type
Select Next: Permissions to move to the next step
Select Attach existing policies directly
In the Search box next to Filter policies, enter ReadOnlyAccess
Scroll to the very bottom of the Policy list until you see the ReadOnlyAccess policy then select the ReadOnlyAccess policy
Select Next: Tags
Select Next: Review
You should see a list containing the trustero-api-user, copy the Access key ID and the Secret access key into the form below.
If you already have an AWS user designated to make read only API calls, you can add an access key. Follow these steps to add an access key:

Log into the AWS IAM console
Select Users from the vertical menu column on the left side of the screen
Select the designated user with ReadOnlyAccess policy privilege
Select Create access key
Use the Access key ID and Secret access key to connect the AWS Receptor to your account
```

### Required Functions

As a developer, you will need to implement the following functions to have a working receptor:
1. func (r *Receptor) GetReceptorType() string {}
    - This function will return a string, signifying the receptor type (e.g. “trr-gitlab”)
2. func (r *Receptor) GetKnownServices() []string {}
    - This function will return an array of string, signifying a list of service types the receptor collects
3. func (r *Receptor) GetCredentialObj() (credentialObj interface{}) {}
4. func (r *Receptor) Verify(credentials interface{}) (ok bool, err error) {}
5. func (r *Receptor) Discover(credentials interface{}) (svcs []*receptor_v1.ServiceEntity, err error) {}
6. func (r *Receptor) Report(credentials interface{}) (evidences []*receptor_sdk.Evidence, err error) {}

#### Get Receptor Type

Returns the “type” of the receptor, usually the name

Ex: “trr-aws”

#### Get Known Services

Returns a list of services that this receptor gathers information for.

Ex: “S3, RDS, IAM”

#### Get Credential Obj

Returns a struct with the needed credential fields
Ex: 
```
{
    "credentials": [
      {
        "display": "Access key ID",
        "placeholder": "ID",
        "field": "akid"
      },
      {
        "display": "Secret access key",
        "placeholder": "Secret access key",
        "field": "skey"
      }
    ]
}
```
The `display` denotes what will be displayed in the UI above the text input box.

The `placeholder` denotes what will be displayed inside the text input box when it is in empty state.

The `field` is the name of the field when the credentials are used in the receptor. This is not visible in the UI.

#### Verify

Very checks the credentials.

The Verify function takes in a set of credentials and makes an API call to the service to confirm the credentials are valid.

It returns a boolean value of “true” if the credentials are valid, or “false” if they are not valid.

#### Discover

Discover finds what is there to learn about.

The Discover function returns a list of Service Entities. This function makes any relevant API calls to the Service Provider to gather information about how many Service Entity Instances are in use (e.g. “How many IAM users?”).

Taking S3 as an example, the Discover call queries AWS for a list of all S3 buckets. That list of S3 buckets would then be converted into a list of `receptor_v1.ServiceEntity` structs, and returned.

#### Report

Report gets details and structures them.

The Report function returns a list of Evidences, each with one or more rows of data. Report makes the needed API calls to gather all data it needs, then generate evidence objects.

This is done with the help of the `receptor_sdk.NewEvidence` function.

The `receptor_sdk.NewEvidence` function takes in :
- Service name, e.g. “S3”
- Entity type, e.g. “bucket”
- Caption, e.g. “S3 Bucket inventory”
- Description, e.g. “A list of S3 buckets with user data”

This will create an `Evidence` object. Individual S3 buckets can then be added to this evidence object via the `evidence.AddRow` call.

The `evidence.AddRow` call takes in an interface of the data that is to be displayed. 

For example, the S3 evidence should show a table with a list of all S3 buckets, when they were created, and their encryption status.
-- --
**_IMPORTANT - You will need to keep track of all API calls made to gather evidence. Each evidence object should also include the raw API call made, as well as the raw JSON response from the api call_**
-- --
To add the raw API call and raw JSON response to the evidence object, you wil need to use the “Evidence.AddSource” function

For S3, the Evidence object would look like this:

```
{
 "caption": "AWS S3 Bucket Encryption Status",
 "description": "A list of AWS S3 buckets and their encryption settings",
 "service_name": "S3",
 "entity_type": "bucket",
 "sources": [
   {
     "raw_api_request": "s3.getAllBuckets",
     "raw_api_response": "{buckets: [Logo-storage-bucket, …]}"
   }
 ],
 "evidence_type": {
   "struct": {
     "rows": [
       {
         "bucket": { "name": "Logo-storage-bucket", "encryption": "AES-256" }
       }
     ]
   }
 }
}

```

Each Service Entity should be represented by one Evidence, with the Service Entity Instances represented as rows within the Evidence. So a list of S3 buckets would be one evidence, and a list of ECS instances would be a different evidence.

## Testing A Receptor

You should be able to run your receptor code via the command line to confirm the Verify and Scan functions produce the correct output.
You can run the Receptor code with the `dryrun` flag and it will print the output to the console.
You can compile the Receptor code into a binary or run the main file directly.
If you run the main file directly, your command should look something like this:

```
go run main.go scan dryrun --find-evidence
```

This command will run the Verify, Discover, and Report functions that you wrote and print their output to the console. You should be able to see the final Evidences that are generated by the receptor.
