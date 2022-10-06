# Help on configuration

This program lets you define the location of the configuration file with the command line argument
`cfgpath` else it will default to `etc/[environment].env` where environment is another argument 
passed through the flag `env`, defaults to *local*

The program panics if it can't find a `.env` file in any of the paths

## The env file variables

The env files are ignored by this repository git configuration the reason 
is because most of its data is sensitive and private

```shell
export MONGODB_CERT=/var/certs/atlas.mongodb.cloud.generated.certificate.file.pem
# MONGODB_CERT is not needed if you connect through username/password or pass it directly
# to the URL
export MONGODB_URI=mongodb+srv://devtest.your.mongodb.cloud.instance.net/?authSource=%24external\
&authMechanism=MONGODB-X509&retryWrites=true&w=majority&tlsCertificateKeyFile=$MONGODB_CERT
# If connect through MONGODB-X509, yiu can use $MONGODB_CERT variable in the following fashion
# &tlsCertificateKeyFile=$MONGODB_CERT
# Connection 

#Google OAuth Client ID and Secret can be obtained from your Google Cloud Console
export GOOGLE_OAUTH_CLIENT_ID=**********************************
export GOOGLE_OAUTH_CLIENT_SECRET=*******************************
# Session secret is used to encrypt your cookie data, ideally it should be a random long string
# to mitigate the risk of dictionary attacks something like the one below
export SESSION_SECRET=d5zxs9CV4UqAAHvKiX26mg
```
