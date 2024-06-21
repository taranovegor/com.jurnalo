# ĵurnalo!
Collect and monitor logs

>[!Warning]
>**ĵurnalo is currently in Development**: Expect breaking changes and bugs!

## System Requirements
- Architectures: arm64, amd64
- Operating Systems: Linux Debian-based
- Database: MongoDB 3.6 or higher
- Software: systemd installed

## Configuration File

The application requires a configuration file to set up necessary parameters. Below is an example of the configuration file `/etc/jurnalo/config.yaml`:
```yaml
http:
    port: 5341 # the HTTP port on which the web interface will be accessible
database:
    uri: mongodb://localhost:27017 # the URI for connecting to the MongoDB database
    name: jurnalo # the name of the MongoDB database to connect to
```
