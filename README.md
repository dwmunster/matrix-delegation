Matrix Delegation Handler
=========================

This package serves to handle the `/.well-known/matrix/{server,client}` URLs as defined in the [Matrix server spec](https://matrix.org/docs/spec/server_server/latest#get-well-known-matrix-server).


Arguments
---------
The following arguments (or environment variables) can be provided.

- `-homeserver` (`HOMESERVER_ADDRESS`): the address of the homeserver to which the domain delegates for federation. 
  This argument/variable is mandatory.
- `-baseurl` (`BASE_URL`): the base url of the homeserver from a client's perspective. This argument/variable is optional, but the application will not server `/.well-known/matrix/client` without it.
- `-listen` (`LISTEN_ADDRESS`): the address on which to listen. Optional, defaults to `:8000`.

Endpoints
---------
- `/.well-known/matrix/server`
- `/.well-known/matrix/client`
- `/health` (returns `200` if running)