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

### Post

#### ✅ \[POST\] /api/post \[AUTH Required\]

Description: Posts a new post with the body of "title, body, forumName" using the auth token.

Request Fields: title, body, forumName

Response on Success: JSON object with "id" int64. Should be managed as a string on platforms and languages without a 64 bit integer type.

#### ✅ \[GET\] /api/post/<forum_name>/<post_id> \[AUTH OPTIONAL\]

Description: Get Post with Id

Work: Status Done.

#### ✅ \[POST\] /api/post/:forumName/:id/upvote \[AUTH REQUIRED\]

Description: Adds upvote. changes vote if already present

Work: Status Done

#### ✅ \[POST\] /api/post/:forumName/:id/downvote \[AUTH REQUIRED\]

Description: Adds down vote. changes vote if already present.

Work: Status Done

#### ✅ \[DELETE\] /api/post/:forumName/:id/deletevote \[AUTH REQUIRED\]

Description: removes vote if present. otherwise, silently fails

Work: Status Done

### Forum

#### ✅ \[POST\] /api/forum \[AUTH Required\]

Description: Create a new Forum by provide a json body with title and description strings.

Work Status: Done

#### ✅ \[GET\] /api/forum 

Description: Retrieve list of public forums. If authenticated, the list will include private forums you joined / owned.

Work Status: Done but not paginated. Work in the future for pagination.

#### ✅ \[GET\] /api/forum/<forum_name>

Description: Retrieves the forum + posts for the given name

Work Status: Done. However, does not hide private forums.

### User 

#### ✅ \[POST\] /api/user

Description: Creates a new User.

Work Status: Creates the User

#### \[Get\] /api/user/<username>

Description: Gets User Data. If the requester is authenticated, potentially returns more data for self.

Work Status: Not Started

### Session

#### ✅ \[POST\] /api/session/new 

body(forum|json): username & password

return data: Bearer-Token

Description: Creates a new session using username password

Work Status: Endpoint Functional

#### ✅ \[Delete\] /api/session

header: Bearer-Token

body: none

Work Status: Endpoint Functional

#### \[POST\] /api/session

header: Bearer-Token

body: none

return data: New Bearer-Token

Work Status: Not Started

### Group

Description: Manages a group of posts that belong to 1 or more unlinked forums

#### ✅ \[GET\] /api/group \[AUTH Optiona\]

return data: posts

Work Status: Returns everything for now.


#### Event Websocket

When points of interest change, the rest api will push alerts (and potential updates).

### Storage

The current designated storage backend for the backend service is Postgres.
