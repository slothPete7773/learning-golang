# Learn Pubsub system with Emitter.io

## Setup

0. **.env file**

    ```sh
    $ cp .env.sample .env
    ```

1. **Instantiate Emitter server.**

    For first time setup, we won't have a License key for environment variable. Run the server 1 time and it will generate for us in the Docker CLI.

    ```sh
    $ docker compose up
    ```


2. **Configure License environment variable.**

   - Copy the License to `.env`'s `EMITTER_LICENSE` variable.
   - Copy the secret key in the `.env`'s `EMITTER_SECRET_KEY`. This secret key is use for generate a channel key in the `localhost:[port\]/keygen`

3. **Instantiate Emitter server again.**

    ```sh
    $ docker compose up
    ```


## Reference
- https://hub.docker.com/r/emitter/server/
- https://emitter.io/develop/what-is-emitter/
- https://github.com/emitter-io/go
- https://emitter.io/download/