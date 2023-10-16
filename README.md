# Rename Forums

- Forums Software

## Architecture

Currently, there are two main parts: the frontend app ( Website ) and the backend service (webapi)

### The Website

The website is a server side rendered (hydrated) web app. Its entire purpose is to provide a user
accessible access point to the backend service. It should be replacable at a glance.

### The backend service

The intended purpose of the backend is to provide an easy way to pull and push data from clients
and storage. Information pushed will be stored and events will be pushed to clients to indicate
the client should refresh. An optional live feed may be desirable in the future.

#### REST API

The REST api will provide access to all information stored in storage. 

##### Gets all forums list

/api/forums?page=</number/>&by=</number/>

##### Gets forum home posts

/api/forum/</name/>?page=>&by=</number/>

##### Gets a post of a forum

/api/forum/</name/>/</postname/>?comment_length=</number/>&comment_depth=</number/>&max_comment=</number/>

##### Create User/POST

/api/user/create

##### Create Session/POST ✅

Work Status: Endpoint created

/api/session/new

##### Delete Session/DELETE ✅

Work Status: Endpoint Functional

/api/session

##### Refresh Session/POST

/api/session

##### Get User/GET

/api/user/username

#### Event Websocket

When points of interest change, the rest api will push alerts (and potential updates).

### Storage

The current designated storage backend for the backend service is Postgres.
