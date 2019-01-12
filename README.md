# Welcome to Yawoen company data API

Here you will find the source code, manual to deploying and testing the application.

**Technologies**: [Golang](https://golang.org/), [Google Cloud App Engine](https://cloud.google.com/appengine/) and [Google Cloud Datastore](https://cloud.google.com/datastore/).

The APIs are running in the App Engine, and recording the data in Datastore database.

There are three API's:

 - To load initial data: https://neoway.appspot.com/load
 - To integrate data: https://neoway.appspot.com/integrate
 - To get data: https://neoway.appspot.com/consult
 

## Deploying

The application was deployed on Google Cloud Platform. Follow the steps to deploy your own application:

1) Install [Golang](https://golang.org/doc/install).

2) Create [Google Cloud](http://cloud.google.com) account. Go to GCP and create a new free account, they give $300 and 12 months period to use the service without payment.
    
3) In the [console](https://console.cloud.google.com) create a new project.

4) Install [Google Cloud SDK](https://cloud.google.com/sdk/docs/quickstarts).

5) Git source code to GOPATH folder in your pc, into the src folder. The folder structure must be /src/projetos/neoway. Important pay attention to create "projetos" folder.

6) Edit the file app.yaml , put your project code (line 20). 

**![](https://lh4.googleusercontent.com/LgjfPdhwgagCIIoFRlS1lq_sbo81luvyVbljjHn2XaAr0-fVgDdwfKxehXzYLzefZ2if2M59NUGzWNREz-fHtm1wGRTg4Lph2wxaOTD58IHIjNRsmF_SK5Ddga7VXyrcusUr5ap1)**

**![](https://lh5.googleusercontent.com/qCbpYJId1dShaL-XMizSgaMoYBPFPqkziKwIRtplfwapMHjBJ9RUPG6-OQg_Ee8KB2823QTG3Hi5bO5DOC-qN9dAVLfJPgCP_uCgAsVI4kvAo_oeZ5vH0_6AYAcUfGPqWkOg-ASb)**

7) The following steps are performed on the terminal with the Google Cloud SDK (gcloud).

	1 - $ **gcloud auth login** (To authorize login in your google account).
	
	2 - $ **gloud config set project PROJECTCODE** (example: gcloud config set project neoway)
	
	3 - Navigate to the folder that you downloaded the files, check if you are in the folder that contains the files app.yaml and main.go.
	
	4 - $ **gcloud app deploy app.yaml** (to deploy the application)
 
 Choose the region. Suggestion southamerica-east1 (9).
 
 Type y to continue.
 
**![](https://lh3.googleusercontent.com/qOtwhrQJvsqHOv_tePvgWoevrTMbYOKNB8d5jdy11c7keDUhzllY8IbxnjoXMsR0d9-arYqsWBTvvg9BikDXImT-4n0n7GwQuBGkZCBx3yP-VErBffilvnFEgjQVQVaDUYWs1wQt)**

**Wait a few minutes….**

When the process finish you will see that.
**![](https://lh4.googleusercontent.com/eksx70j_Mp1XDamOgfLrM3d9_-iLrI4N8M6H6lvxgDeqAMQcU5jrVwGmzRRElpXn5LvgIdnysZDjovDXoNsKejs5s90s-Velg2SpDAAKJTp7P5Z9UlwYeqv7JjpouyAkS0lDZKzr)**

**OK, it's done.**

## Testing

