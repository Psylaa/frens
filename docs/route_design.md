# This is very, very much a work in progress

The process of designing API routes is an iterative and often challenging one, frequently involving revisiting assumptions, addressing overlooked use cases, and reconfiguring organization. This can lead to a time-consuming cycle of constant modifications, spec rewriting, and API redesigns that slow down the overall development process and hinder efficiency.

In an effort to streamline this process, we have adopted an approach focused on detailed use case analysis. The goal of this approach is to identify and define every possible practical use case for each database item, essentially answering the question: what actions might a user need to perform with this object?

By thoroughly examining each use case, we can ensure that our API routes are comprehensive, and each use case is addressed and covered adequately. This preemptive strategy aims to minimize the need for major API redesigns later on, thereby saving time and effort. Once the analysis is complete, we use it as a robust guide to design our API, allowing us to confidently cover most, if not all, anticipated use cases.

This document, therefore, serves as an integral part of our API design process. By laying out all the potential interactions between a user and the objects in our database, it provides us with a roadmap to design a comprehensive and efficient API system that confidently addresses the needs of our users. If believe we missed a use case, please feel free to open an issue.

### Sorting
Sort each item in order of method (Post > Get > Patch > Put > Delete), route length (/thing > /thing/thing2)

# Blocks

Block - A source user ID and target user ID. 

| Use Case                            | HTTP Method | Path                     | Status  |
| ----------------------------------- | ----------- | ------------------------ | ------- |
| Block a user                        | POST        | /users/{usersID}/blocks  | Pending |
| Get users I've blocked              | GET         | /blocks                  | Pending |
| Get a block by ID                   | GET         | /blocks/{blockID}        | Pending |
| Get users I've blocked (pagination) | GET         | /blocks?count=n&offset=m | Pending |
| Remove all blocks                   | DELETE      | /blocks                  | Pending |
| Unblock a user                      | DELETE      | /users/{userID}/blocks   | Pending |

# Bookmarks

Bookmark - A user ID and post ID.

| Use Case                                     | HTTP Method | Path                            | Status  |
| -------------------------------------------- | ----------- | ------------------------------- | ------- |
| Create a bookmark                            | POST        | /posts/{postID}/bookmarks       | Covered |
| View all my bookmarks                        | GET         | /bookmarks                      | Covered |
| View my bookmarks in batches (pagination)    | GET         | /bookmarks?count=n&offset=m     | Covered |
| View one of my bookmarks by bookmark ID      | GET         | /bookmarks/{bookmarkID}         | Pending |
| View one of my bookmarks by post ID          | GET         | /posts/{postID}/bookmarks       | Pending |
| View the total count of bookmarks for a post | GET         | /posts/{postID}/bookmarks       | Pending |
| Check if I've bookmarked a post              | GET         | /posts/{postID}/bookmarks       | Pending |
| View the number of bookmarks a post has      | GET         | /posts/{postID}/bookmarks/count | Pending |
| Delete a bookmark by bookmark ID             | DELETE      | /bookmarks/{bookmarkID}         | Covered |
| Delete a bookmark by post ID                 | DELETE      | /posts/{postID}/bookmarks       | Pending |
# Follows

Follow - A source user ID and target user ID.

| Use Case                                                    | HTTP Method | Path                                       | Status  |
| ----------------------------------------------------------- | ----------- | ------------------------------------------ | ------- |
| Follow a user                                               | POST        | /users/{userID}/follows                    | Pending |
| View all users I am following (pagination)                  | GET         | /users/{userID}/follows?count=n&offest=m   | Pending |
| View all users that someone else is following (pagination)  | GET         | /users/{userID}/follows?count=n&offset=m   | Pending |
| View all users that are following me (pagination)           | GET         | /users/{userID}/followers?count=n&offset=m | Pending |
| View all users that are following someone else (pagination) | GET         | /users/{userID}/followers?count=n&offset=m | Pending |
| Unfollow a user                                             | DELETE      | /users/{userID}/follows                    | Pending |

# Files

File - A reference to a file on disk with extension and owner (user ID). 

| Use Case                                       | HTTP Method | Path                    | Status  |
| ---------------------------------------------- | ----------- | ----------------------- | ------- |
| Upload a new file                              | POST        | /files                  | Pending |
| Get all the files I have uploaded              | GET         | /files                  | Pending |
| Get all the files I have uploaded (pagination) | GET         | /files?count=n&offset=m | Pending |
| Get a file by file ID                          | GET         | /files/{fileID}         | Pending |
| Delete a file by file ID                       | DELETE      | /files/{fileID}         | Pending |

# Likes

Like - A userID and postID.

| Use Case                                          | HTTP Method | Path                                 | Status  |
| ------------------------------------------------- | ----------- | ------------------------------------ | ------- |
| Like a post                                       | POST        | /posts/{postID}/likes                | Pending |
| View all likes I have given                       | GET         | /likes                               | Pending |
| View all posts I have liked                       | GET         | /likes                               | Pending |
| View a specific like by like ID                   | GET         | /likes/{likeID}                      | Pending |
| Check if a specific like exists by like ID        | GET         | /likes/{likeID}                      | Pending |
| View a specific like by post ID                   | GET         | /posts/{postID}/likes                | Pending |
| Check if a post has been liked                    | GET         | /posts/{postID}/likes                | Pending |
| View all likes for a specific post                | GET         | /posts/{postID}/likes                | Pending |
| View the number of likes a post has               | GET         | /posts/{postID}/likes/count          | Pending |
| View the latest posts I have liked                | GET         | /likes?sort=latest                   | Pending |
| View my likes in batches (pagination)             | GET         | /likes?count=n&offset=m              | Pending |
| Check if a specific like exists by post ID        | GET         | /posts/{postID}/likes                | Pending |
| View all likes from a specific user               | GET         | /users/{userID}/likes                | Pending |
| View the number of likes a specific user has made | GET         | /users/{userID}/likes/count          | Pending |
| Check if a user has liked a specific post         | GET         | /users/{userID}/posts/{postID}/likes | Pending |
| View all users who have liked a specific post     | GET         | /posts/{postID}/likes/users          | Pending |
| Unlike a post by like ID                          | DELETE      | /likes/{likeID}                      | Pending |
| Unlike a post by post ID                          | DELETE      | /posts/{postID}/likes                | Pending |

# Messages
| Use Case | HTTP Method | Path | Status |
| -------- | ----------- | ---- | ------ |

To be implemented at a later date.

# Notifications
| Use Case                          | HTTP Method | Path                            | Status  |
| --------------------------------- | ----------- | ------------------------------- | ------- |
| Get my notifications              | GET         | /notifications                  | Pending |
| Get my notifications (pagination) | GET         | /notifications?count=n&offset=m | Pending |
| Mark all notifications as read    | PUT         | /notifications                  | Pending |
| Mark a notification as read       | PUT         | /notifications/{notificationID} | Pending |
| Mark a nofitication as unread     | PUT         | /notifications/{notificationID} | Pending |

# Posts

| Use Case                                                     | HTTP Method | Path                                    | Status  |
| ------------------------------------------------------------ | ----------- | --------------------------------------- | ------- |
| Create a new post                                            | POST        | /posts                                  | Pending |
| View a specific post by post ID                              | GET         | /posts/{postID}                         | Pending |
| Update a post by post ID                                     | PUT         | /posts/{postID}                         | Pending |
| Delete a post by post ID                                     | DELETE      | /posts/{postID}                         | Pending |
| View all my posts                                            | GET         | /users/{userID}/posts                   | Pending |
| View all posts by a specific user                            | GET         | /users/{userID}/posts                   | Pending |
| Attach a file to a post                                      | POST        | /posts/{postID}/files/{fileID}          | Pending |
| View all files attached to a post                            | GET         | /posts/{postID}/files                   | Pending |
| Remove a file from a post                                    | DELETE      | /posts/{postID}/files/{fileID}          | Pending |
| View the number of likes a post has                          | GET         | /posts/{postID}/likes/count             | Pending |
| View all my posts in batches (pagination)                    | GET         | /users/{userID}/posts?count=n&offset=m  | Pending |
| View all posts by a specific user in batches (pagination)    | GET         | /users/{userID}/posts?count=n&offset=m  | Pending |
| Share a post                                                 | POST        | /posts/{postID}/shares                  | Pending |
| View the number of shares a post has                         | GET         | /posts/{postID}/shares/count            | Pending |
| View all posts I shared                                      | GET         | /users/{userID}/shares                  | Pending |
| View all posts I shared in batches (pagination)              | GET         | /users/{userID}/shares?count=n&offset=m | Pending |
| Set the privacy of a post                                    | PUT         | /posts/{postID}/privacy                 | Pending |
| View all posts by privacy setting (public, private, friends) | GET         | /users/{userID}/posts?privacy={setting} | Pending |
| View the author of a post                                    | GET         | /posts/{postID}/author                  | Pending |
| Get all replies to a post                                    | GET         | /posts/{postID}/replies                 | Pending |

# Reports
| Use Case                     | HTTP Method | Path                    | Status  |
| ---------------------------- | ----------- | ----------------------- | ------- |
| File a report for a user     | POST        | /users/{userID}/reports | Pending |
| File a report for a post     | POST        | /posts/{postID}/reports | Pending |
| View my reports              | GET         | /reports                | Pending |
| View my reports (pagination) | GET         | /reports                | Pending |
| Delete a report              | DELETE      | /reports/{reportID}     | Pending |

# Users

| Use Case                                  | HTTP Method | Path                                           | Status  |
| ----------------------------------------- | ----------- | ---------------------------------------------- | ------- |
| Register a new user                       | POST        | /users                                         | Pending |
| View a specific user by user ID           | GET         | /users/{userID}                                | Pending |
| Update my user profile                    | PUT         | /users/{userID}                                | Pending |
| Delete my user profile                    | DELETE      | /users/{userID}                                | Pending |
| View all my posts                         | GET         | /users/{userID}/posts                          | Pending |
| View all posts of a specific user         | GET         | /users/{userID}/posts                          | Pending |
| View all posts I liked                    | GET         | /users/{userID}/likes                          | Pending |
| View all posts I shared                   | GET         | /users/{userID}/shares                         | Pending |
| View all my bookmarks                     | GET         | /users/{userID}/bookmarks                      | Pending |
| Follow a user                             | POST        | /users/{userID}/follow                         | Pending |
| Unfollow a user                           | DELETE      | /users/{userID}/unfollow                       | Pending |
| View all users I follow                   | GET         | /users/{userID}/follows                        | Pending |
| View all users who follow me              | GET         | /users/{userID}/followers                      | Pending |
| View all users I blocked                  | GET         | /users/{userID}/blocks                         | Pending |
| Change my password                        | PUT         | /users/{userID}/password                       | Pending |
| Reset my password                         | PUT         | /users/{userID}/password/reset                 | Pending |
| Verify my email                           | PUT         | /users/{userID}/email/verify                   | Pending |
| Resend verification email                 | POST        | /users/{userID}/email/verification/resend      | Pending |
| Change my email                           | PUT         | /users/{userID}/email                          | Pending |
| Upload a profile picture                  | POST        | /users/{userID}/profile_picture                | Pending |
| Remove my profile picture                 | DELETE      | /users/{userID}/profile_picture                | Pending |
| Upload a cover picture                    | POST        | /users/{userID}/cover_picture                  | Pending |
| Remove my cover picture                   | DELETE      | /users/{userID}/cover_picture                  | Pending |
| Delete a notification                     | DELETE      | /users/{userID}/notifications/{notificationID} | Pending |
| View all my messages                      | GET         | /users/{userID}/messages                       | Pending |
| Send a message                            | POST        | /users/{userID}/messages                       | Pending |
| Delete a message                          | DELETE      | /users/{userID}/messages/{messageID}           | Pending |
| Mark a message as read                    | PUT         | /users/{userID}/messages/{messageID}           | Pending |
| Block a user                              | POST        | /users/{userID}/blocks/{blockID}               | Pending |
| Unblock a user                            | DELETE      | /users/{userID}/blocks/{blockID}               | Pending |
| View all users muted by a specific user   | GET         | /users/{userID}/mutes                          | Pending |
| View whether a specific user has muted me | GET         | /users/{userID}/mutes/me                       | Pending |
| Mute a user                               | POST        | /users/{userID}/mutes/{muteID}                 | Pending |
| Unmute a user                             | DELETE      | /users/{userID}/mutes/{muteID}                 | Pending |
| Update my profile picture                 | PUT         | /users/{userID}/profile_picture                | Pending |
| Remove my profile picture                 | DELETE      | /users/{userID}/profile_picture                | Pending |
| Update my cover photo                     | PUT         | /users/{userID}/cover_photo                    | Pending |
| Remove my cover photo                     | DELETE      | /users/{userID}/cover_photo                    | Pending |
| Update my bio                             | PUT         | /users/{userID}/bio                            | Pending |
| Remove my bio                             | DELETE      | /users/{userID}/bio                            | Pending |
| View user's profile picture               | GET         | /users/{userID}                                | Pending |
| View user's cover photo                   | GET         | /users/{userID}                                | Pending |
| View user's bio                           | GET         | /users/{userID}                                | Pending |
| Update my password                        | PUT         | /users/{userID}/password                       | Pending |
| Set a file as profile picture             | POST        | /users/{ownerID}/profilePicture/{fileID}       | Pending |
| View a user's profile picture             | GET         | /users/{userID}/profilePicture                 | Pending |
| Remove a profile picture                  | DELETE      | /users/{ownerID}/profilePicture/{fileID}       | Pending |
| Set a file as cover image                 | POST        | /users/{ownerID}/coverImage/{fileID}           | Pending |
| View a user's cover image                 | GET         | /users/{userID}/coverImage                     | Pending |
| Remove a cover image                      | DELETE      | /users/{ownerID}/coverImage/{fileID}           | Pending |