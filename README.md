# Server Scripts

This is a collection of application maintenance scripts
to be run on the server.  The scripts expect the following
directory structure:

- /var/www/application
- /var/www/application_conf
- /var/www/application_logs

Replace _application_ with the name of your application.

Production configuration and logging configuration must be stored in 
application_conf and named as follows:

- /var/www/application_conf/production.conf
- /var/www/application_logs/logback.xml

### Setup

Create production directories:

- /var/www/application
- /var/www/application_conf
- /var/www/application_logs

Unzips the application package, installs and runs the application.

### Go

Starts MongoDB, then starts the application. Run this script after server reboot.

### Start

Starts the application.

### Stop

Stops the application.

### Update

Updates the application to a new version.  Requires ZIP file of the new version to be 
in the same directory as the script.

## Configuring Scripts

Modify APP_DIR and APP_DIR_CONF variables in each script to reflect the name
of your application.  For example:

- APP_DIR = coolapp
- APP_DIR_CONF = coolapp_conf

## Running Scripts

Upload the scripts to your user directory on the server. Login into the server
and set permission to the script files to read, write, and execute.

- `sudo chmod 700 go`
- `sudo chmod 700 start`
- `sudo chmod 700 stop`
- `sudo chmod 700 update`

Execute script with **sudo**:

`sudo ./go`

