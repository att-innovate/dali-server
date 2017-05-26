## Dali Server API

Following APIs are offered by the dali-server. `127.0.0.1` has to be replaced with the IP address of your dali-server.

### /admin/reset

Resets the database. Clear all the marks.

Example curl request to local server to wipe database:

	curl -X GET -H "Accept: application/json" -H "Content-Type: application/json" 127.0.0.1:8085/admin/reset

Example response:
	
	{"code":202,"message":"done"}

### /id

Returns a new unique id needed to create mark.

Example curl GET request to get a new unique id:

	curl -X GET -H "Accept: application/json" -H "Content-Type: application/json" 127.0.0.1:8085/id

Example response from server:
	
	{"code":202,"message":"57c9c16868116a000197800c"}


### /mark 

**POST** ---> posts a new mark to the mark database

Example curl POST request to local server:

	curl -X POST -H "Accept: application/json" -H "Content-Type: application/json" 127.0.0.1:8085/mark/ -d '{"id":"57a3d4cee7256911701e1674", "Label":"Front Door","Type":"tasklist","Content":["Turn off lights","Activate Alarm"]}'

Example response from server:

	{"code":202,"message":"57a3d4cee7256911701e1674"}

**GET** ---> returns all the marks with all the details

Example curl GET request to local server:

	curl -X GET -H "Accept: application/json" -H "Content-Type: application/json" 127.0.0.1:8085/mark/

Example response to GET request to local server:

	{"marks":[{"Id":"57a3d4cee7256911701e1674","Label":"Front Door","Type":"tasklist","Content":["Turn off lights","Activate Alarm"]}]}


### /template/(file name)

Renders a mark using the type specific template.

Example URL to view a specific mark in your web browser:

[http://127.0.0.1:8085/template/tasklist.html?id=57a3d4cee7256911701e1674](http://127.0.0.1:8085/template/tasklist.html?id=57a3d4cee7256911701e1674)

`taslklist.html` is the html template which needs to match the type of annotation of the mark defined by the `id`.


### /mark/(id number)
----------------

**GET** ---> returns the mark associated with the specific id

Example curl GET request to retrieve mark:

	curl -X GET -H "Accept: application/json" -H "Content-Type: application/json" 127.0.0.1:8085/mark/57a3d4cee7256911701e1674

Example response from server:

	{"Id":"57a3d4cee7256911701e1674","Label":"Front Door","Type":"tasklist","Content":["Turn off lights","Activate Alarm"]}
 
**PUT** ---> Update the mark associate with the specified id

Example curl PUT request to local server:

	curl -X PUT -H "Accept: application/json" -H "Content-Type: application/json" 127.0.0.1:8085/mark/57a3d4cee7256911701e1674 -d '{"id":"57a3d4cee7256911701e1674", "Label": "Exit", "Type":"image", "Content":"UpArrow"}'

Example response from server:

	{"code":202,"message":"done"}

**DELETE** --> Deletes mark associated with this id

Example curl DELETE request to local server:

	curl -X DELETE -H "Accept: application/json" -H "Content-Type: application/json" 127.0.0.1:8085/mark/57a3d4cee7256911701e1674 

Example response from server:

	{"code":202,"message":"done"}


