# help-people-api
People Helping People API Server

## API

All proceeding routes should be prefixed by the api version number (`/v1/needs`).

`*` indicates the route does not require an auth token

### need-types
  - `*GET need-types/` - Fetches an array of available types of needs that can be created

### needs
  - `GET needs/` - Gets available needs in the helpers local area
  - `GET needs/{id}` - Fetches info about the need
  - `*POST needs/{id}` - Creates a new need in the system and returns an updated JWT token that grants access/editing of the need
    - `needType` String - The type of need this is
    - `name` String (optional) - An optional name to include with this request
    - `description` String (optional) - A brief description or details for the need
  - `PATCH needs/{id}` - Updates the need request using the data in the given body
    - `name` String (optional) - An optional name to include with this request
    - `description` String (optional) - A brief description or details for the need
  - `DELETE needs/{id}` - Archives the given need (this can only be archived using the special need token)
  - `GET needs/{id}/helpers` - Gets all helpers willing to assist in this need
  - `POST needs/{id}/helpers/{hid}` - Marks a helpers offer of assistance as accepted
  - `DELETE needs/{id}/helpers/{hid}` - Marks a helpers offer of assistance as declined
  - `GET needs/{id}/messages` - Gets all messages for this need
  - `POST needs/{id}/messages` - Creates/sends a new message for this need


### helpers
  - `GET helpers/{id}` - fetches the helpers info
    - `firstName` String - users first name
    - `lastName` String - users last name
    - `supports` Array - List of needs the user is willing to fulfill
  - `*POST helpers/` - creates a new helper account with given info in the body
    - `firstName` String - users first name
    - `lastName` String - users last name
    - `password` String - users password
    - `supports` Array - List of needs the user is willing to fulfill
    - `deviceToken` String - The unique device token
  - `PATCH helpers/{id}` - updates the user profile with info in body
    - `firstName` String - users first name
    - `lastName` String - users last name
    - `password` String - users password
    - `supports` Array - List of needs the user is willing to fulfill
    - `deviceToken` String - The unique device token
  - `DELETE helpers/{id}` - marks the user's profile as inactive
  - `POST helpers/{id}/needs/{nid}` - Marks a need as accepted by the helper
    - `message` String (optional) - An optional message the helper can send when wanting to provide help to a need
  - `DELETE helpers/{id}/needs/{nid}` - Marks a need as dismissed by the helper

### location

  - `POST location/` - Updates the authenticated users location in the system
    - `latitude` - The users new latitudinal coordinates
    - `longitude` - The users new longitudinal coordinates
