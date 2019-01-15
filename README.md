# Welcome to Yawoen company data API

Here you will find the source code, manual to deploying and testing the application.

**Technologies**: [Golang](https://golang.org/), [Google Cloud App Engine](https://cloud.google.com/appengine/) and [Google Cloud Datastore](https://cloud.google.com/datastore/).

The APIs are running in the App Engine, and recording the data in Datastore database.

There are three API's:

 - To load data: https://neoway.appspot.com/load
 - To integrate data: https://neoway.appspot.com/integrate
 - To get data: https://neoway.appspot.com/consult/{name}/{zip}

## API Roles

### API Load data
This api creates new entities with name and zip code in the database.

**Data format:**
- Apply upper to the name.


### API Integrate data
This api finds the entity in the database using the name.
- If find the entity, record the website.
- If can not find the entity, crate a new company in database.

**Data format:**
- Apply upper to the name.
- Apply lower to the website.

### API Get data
This api finds the entity in the database using part of the name and zip code. 
Returns the json object with id, name, zip and website. 

## Deploying

The application was deployed on Google Cloud Platform. Follow the steps to deploy your own application:

1) Install [Golang](https://golang.org/doc/install) in your pc.

2) Create [Google Cloud](http://cloud.google.com) account. Go to GCP and create a new free account, they give U$ 300 and 12 months period to use the service without payment.
    
3) In the [console](https://console.cloud.google.com) create a new project.

4) Install [Google Cloud SDK](https://cloud.google.com/sdk/docs/quickstarts) in your pc.

5) Git source code to GOPATH folder in your pc, into the src folder. The folder structure must be /src/projetos/neoway. **Important** pay attention to create "projetos" folder.

6) Get the needed packages, executing the file ```goGet.sh```.
```
$ ./goGet.sh
```
7) Edit the file app.yaml , put your project code (line 20). 

**![](https://lh4.googleusercontent.com/LgjfPdhwgagCIIoFRlS1lq_sbo81luvyVbljjHn2XaAr0-fVgDdwfKxehXzYLzefZ2if2M59NUGzWNREz-fHtm1wGRTg4Lph2wxaOTD58IHIjNRsmF_SK5Ddga7VXyrcusUr5ap1)**

**![](https://lh5.googleusercontent.com/qCbpYJId1dShaL-XMizSgaMoYBPFPqkziKwIRtplfwapMHjBJ9RUPG6-OQg_Ee8KB2823QTG3Hi5bO5DOC-qN9dAVLfJPgCP_uCgAsVI4kvAo_oeZ5vH0_6AYAcUfGPqWkOg-ASb)**

8) Finally the following steps are performed on the terminal with the Google Cloud SDK (gcloud).
```
$ gcloud auth login (To authorize login in your google account).
````
```	
$ gloud config set project PROJECTCODE (example: gcloud config set project neoway)
```	
Navigate to the folder that you downloaded the files, check if you are in the folder that contains the files app.yaml and main.go.
```	
$ gcloud app deploy app.yaml (to deploy the application)
```
 	
Choose the region. Suggestion southamerica-east1 (9).
 
Type y to continue.

![](https://lh3.googleusercontent.com/qOtwhrQJvsqHOv_tePvgWoevrTMbYOKNB8d5jdy11c7keDUhzllY8IbxnjoXMsR0d9-arYqsWBTvvg9BikDXImT-4n0n7GwQuBGkZCBx3yP-VErBffilvnFEgjQVQVaDUYWs1wQt)

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

You can test with my urls, i left the database empty. But you don't have permission to check the database in my project. If you send me your google account, i'll give you access to the project (send to e-mail evandropp10@gmail.com).
But you can test with your own url in your own project.

[Postman](https://www.getpostman.com/) software was used, but you can use another similar.

You can import the configuration in Postman. Just click in the Import button, choose "Import from link" and paste this link: https://www.getpostman.com/collections/f955c2e0a28cd9589603

1.  LOAD INITIAL DATA

POST csv file in the load url.

**![](https://lh6.googleusercontent.com/58JxFib61jcsqfQ7aY3ndSGRq8WxOSUGIHfVPVE9en4MECWUnUkTNgq2b3g4fvwCKR76VbTvCerJ3mjgvUtYqgC3U1tY818Q6JtWgfkJOSEcOxZiMTj7TC8DOghX2pBazUfU58xX)**

Return Status:

![](https://lh4.googleusercontent.com/XRAwdmRm2XJ23dzLZXo8zn9vias0N1BL81YTbDkNGSuBBMdPMERV7xMy5LfNIHAFgfUzyrndKhPO4Ti61MhBdoC1YUXnSKFny-oTLdDJZ2QwrbfMWhw7YIPXN31m1U3qZaN5nYFs)

2) CHECK DATABASE

Maybe you'll have to choose the project, and you will see the companies entities with the columns ID/Nome, Name, Zip and Website.

Go to [https://console.cloud.google.com/datastore](https://console.cloud.google.com/datastore) and check the entities on database.

![](https://lh6.googleusercontent.com/rhfQ6WJ6L58NkyybemOaCxl8O07C9GdycPdLFbh0SpjtBjrEzpxk7BaI3bxQ98Qc0VDmpl75AdRGWPpVI6q2n0RnIKmZGtjEJeARsWQGpv0Cm0fhkUK8bq5Z-IUki7NL-YE6PGeY)

3) REQUEST INTEGRATE DATA

POST csv file in the load url.

![](https://lh4.googleusercontent.com/kjzXoRmhsiXtgd5DUMeeoMg0p5ESaJuS-7QoCR5UQyfNYn3mPxPUXRqrvUbrNOWh8odp5JUOqjhEKZAVxlqKBS1ff4QBRBC1nhtsvyWOdgU61YIGrG2R97bPCuT1QXDWdRg0mLl2)

4) CHECK DATABASE

The column Website was populated.

**![](https://lh3.googleusercontent.com/9XIRJNMDQaG9qZpUe3XagzjKiJfGwS-Kd8ABSMGLCjMdqzoMyxLqs6UYUvxFEqvACDpM8mHZVHzkNeW6-C8hgIyHAr5peJcbE1GEqRcbbYbpGAlvK63c3wsQxXph0GJJ-JseshCw)**

5) GET COMPANY

Request the Get method in the api [https://neoway.appspot.com/consult/{name}/{zip](https://neoway.appspot.com/consult/%7Bname%7D/%7Bzip)}. It's mandatory to set the parameters name and zip.

The format of the returned object is:
```
{
	"id":"5195749109268480",
	"name":"CRICKET WIRELESS AUTHORIZED RETAILER",
	"zip":"77009",
	"website":"https://www.cricketwireless.com"	
}
```

**![](https://lh3.googleusercontent.com/uHDWdmxUXagyVwClrBdtB1m3cVRNVhjbnKfci5tybm_Je4tMWRKa4DIw7eqobolDvMXmK0zgVT0YI8L3hXa8392NEL82c_9SV8nN7YE1lrdsXo-nj_Yn5ruQ2d0Od0R5l2weVszr)**

If it does not find in the database, returns not found.
**![](https://lh5.googleusercontent.com/NEPHcvLNxrH10eT1jezOA1hd2Xp7usybe_7X4MMKAPQpUBXfZg26wnWJFdslj59zxWzGjrKDlEf3tbDXsdpXxUkQ4Oe5FLLb_RATwkhptqHws2eUL1GVkSJjGKrO6egZQa6Qs_Cs)**
## Considerations
I chose host in the cloud because i believe that it's a more real and market solution.
Golang, along with Python and Angular, is a language I've been studying since 2015. And I feel comfortable to program with it.

I think that i didn't understand this paragraph well.
"An id field is non existent on the data source, so you'll have to use the available fields to aggregate the new attribute **website** and store it. If the record doesn't exist, discard it."
So i give the following treatment to integrate data:
- If entity exists on database, in this case the entity is modified to add website.
- If entity don't exists, in this case create a new one. I decided don't discard because i think we are losing data doing this.

I hope you enjoy!!
<!--stackedit_data:
eyJoaXN0b3J5IjpbODM2NTY0MzQ5XX0=
-->