# Frens
Welcome to Frens, a decentralized social networking platform engineered for performance, scalability, and respect for user privacy. Our platform is lightweight yet robust, built on Go, and is dedicated to providing an enhanced social media experience that prioritizes user control over data.

## Key Features
One of the biggest blockers to the widespread adoption of decentralized social media is esoteric API standards. Therefore, we have chosen to stricly adhere to the JSON API specification, which is widely used and well-documented. This allows developers to easily integrate Frens into their existing applications, and makes it easy for users to migrate their data to Frens.

Additionally, the API is fully documented with swagger, which allows for easy integration into any application.

## Request Lifecycle
The following section provides a detailed overview of the lifecycle of a request as it's processed by our backend service. This lifecycle comprises of several crucial stages, each playing a unique role in ensuring the successful processing and response to a client's request.

1. Request Reception: The initial request is received by our server which uses the Fiber framework, building on top of the Go's native http package. The role of parsing this request falls to our router package, which is adept at unpacking the incoming request data and reformating it into a structure that our service can readily process.

2. Business Logic Execution: The service package is tasked with processing the parsed request data. It houses the majority of the business logic, including crucial operations such as duplicate checks, data validation, and interaction with the database for persistent storage.

3. Database Interaction: Our database package manages all interactions with the underlying persistent storage layer. It's responsible for all data manipulation tasks including creation, update, and deletion of records, ensuring data consistency and integrity.

4. Response Generation: Finally, the service package coordinates with the response package to generate an appropriate response for the client. The response package is instrumental in ensuring that our responses comply with the JSON API specification, providing a standardized format for the client to consume.

In essence, each stage in this request lifecycle forms an integral part of our service, working together to provide a seamless experience for all users interacting with our backend service.

## Comprehensive Documentation in Our Wiki
Our extensive wiki is your primary resource for all things Frens. Whether you're a user looking to get the most out of the Frens platform, a developer eager to dive into our code, or a curious onlooker who wants to understand our design philosophy and AI integration strategy, the wiki has you covered.

Our wiki is meticulously organized and updated, providing a wealth of knowledge about every aspect of Frens. It serves as a roadmap, guiding you through our:

Design Philosophy: Understand the principles that drive our development process and learn about our AI integration.
API Documentation: Dive into the details of our intuitive API, with clear explanations and examples to ease your development process.
Development Stages: Get a clear idea about our development strategy, from pre-alpha to release stages.
And much more!
Our commitment to transparency and education means our wiki is always growing and evolving, just like Frens itself. It's not just a manual—it's a journey into the heart of our project.

## License
Frens is released under the MIT License.
