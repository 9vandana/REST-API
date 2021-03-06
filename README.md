# REST-API
Develop a RESTful API for CRUD operations on resources using the GoLang and a web toolkit. 
The resources to be managed are objects called sites. A site has the attributes name, role, URL, and a list of zero or more access points (addresses), whereas an access point has the attributes label and URL. All the attributes are strings. Sites are identified by their name, while the access points of a site are identified by their label. Different sites may contain access points with the same label. 
Your API should support CRUD operations on the access points of a site, as well as on the sites. CRUD operations on a site affect its attributes but not its list of access points, with the exception of a delete which cascades to deleting all its access points. You should persist the resources in a plain text local file of json objects. 
Using a database system to persist the resources is optional. Use a RESTful client to test your API. 
Using a json-extension in your web browser and a CLI tool (eg curl or wget) may facilitate the rapid development of a client for testing.
