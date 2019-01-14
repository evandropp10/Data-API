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
My url is [https://neoway.appspot.com](https://neoway.appspot.com). You can see the url in the last screen shot. The url role is [https://YOUR-PROJECT/appspot.com](https://your-project/appspot.com).

The request **header** have two parameters:

- Content-type: application/json
- Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6IiIsInVzZXJuYW1lIjoiIn0.YOyadIOICikZy_0bupJoLDtSWTo2QQA5GKGr-RpIXyQ

**![](https://lh5.googleusercontent.com/y67oEisSkRhnanvUJzFwnxw7WxedE55YvKJEXxCnhE7xiZtqOFh0F-5Eb4fUprRbBB0O6TG4-waJvUmrRazgy4W5RGYLq1W2KNaqO8TKM2MKlEdADGan3Nd3_lF1_1lzX6Q_QM-B)**

You can test with my urls, i left the database empty. But you don't have permission in my project to check the database. If you send to me your google account, i give permission for the project (send to e-mail evandropp10@gmail.com).
But you can test with your own url in your own project.

[Postman](https://www.getpostman.com/) software was used, but you can use another similar.

1.  LOAD INITIAL DATA

POST csv file in the load url.

**![](https://lh6.googleusercontent.com/58JxFib61jcsqfQ7aY3ndSGRq8WxOSUGIHfVPVE9en4MECWUnUkTNgq2b3g4fvwCKR76VbTvCerJ3mjgvUtYqgC3U1tY818Q6JtWgfkJOSEcOxZiMTj7TC8DOghX2pBazUfU58xX)**
Return Status:

![](https://lh4.googleusercontent.com/XRAwdmRm2XJ23dzLZXo8zn9vias0N1BL81YTbDkNGSuBBMdPMERV7xMy5LfNIHAFgfUzyrndKhPO4Ti61MhBdoC1YUXnSKFny-oTLdDJZ2QwrbfMWhw7YIPXN31m1U3qZaN5nYFs)

2) CHECK DATABASE

Maybe you'll have to choose the project, and you will see the companies entities with the columns ID/Nome, Name, Zip and Website.

Go to [https://console.cloud.google.com/datastore](https://console.cloud.google.com/datastore) and check the entities on database.

![](https://lh4.googleusercontent.com/YXjU_v2AnvCxzZEM-AIzUzZsQ9hwH_ySQhAr3fcJgxmV7ISnPZj4ed5bDQ_tLk4RBnKz4foxXjrSU-rwM-1I_6wsUr7CVlDY9wy-LQZ3VQr8ydchBGuoHwy_FmMJqnXH59hEoOfk)

3) REQUEST INTEGRATE DATA.
POST csv file in the load url.

![](https://lh4.googleusercontent.com/kjzXoRmhsiXtgd5DUMeeoMg0p5ESaJuS-7QoCR5UQyfNYn3mPxPUXRqrvUbrNOWh8odp5JUOqjhEKZAVxlqKBS1ff4QBRBC1nhtsvyWOdgU61YIGrG2R97bPCuT1QXDWdRg0mLl2)

4) CHECK DATABASE

The column Website was populated.

![](https://lh6.googleusercontent.com/eZuvl9jCxp6Ophs6tbMEVxUmdIr7TLDSz4RIndWBnpZWsh35HdAZ0ybndERrqWT19kCMessyhx2ajD-woix4N11jCTAwp_Bjea9JchMYkHR0wEzNOZma4wExLxi_RDKS8FPYB1yd)**

<!--stackedit_data:
eyJoaXN0b3J5IjpbLTUwOTk4NjM5MF19
-->