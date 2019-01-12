# Welcome to Yawoen company data API

Here you will find the source code, manual to deploying and testing the application.

**Technologies**: [Golang](https://golang.org/), [Google Cloud App Engine](https://cloud.google.com/appengine/) and [Google Cloud Datastore](https://cloud.google.com/datastore/).

The APIs are running in the App Engine, and record the data in Datastore database.

There are three API's:

 - To load initial data: https://neoway.appspot.com/load
 - To integrate data: https://neoway.appspot.com/integrate
 - To get data: https://neoway.appspot.com/consult
 

# Deploying

The application was deployed on Google Cloud Platform. Follow the steps to deploy your own application:

1) Install [Golang](https://golang.org/doc/install).

2) Create [Google Cloud](http://cloud.google.com) account. Go to GCP and create a new free account, they give $300 and 12 months period to use the service without payment.
    
3) In the [console](https://console.cloud.google.com) create a new project.

4) Install [Google Cloud SDK](https://cloud.google.com/sdk/docs/quickstarts).

5) Git source code to GOPATH folder in your pc, into the src folder. The folder structure must be /src/projetos/neoway. Important pay attention to create "projetos" folder.

6) Edit the file app.yaml , put your project code (line 20). 
