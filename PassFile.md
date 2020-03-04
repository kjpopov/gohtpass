Create the Password File Using the OpenSSL Utilities
If you have OpenSSL installed on your machine, you can create a password file with no additional packages. We will create a hidden file called .htpasswd in the project directory to store our username and password combinations.
You can add a username to the file using this command. We are using sammy as our username, but you can use whatever name you'd like:

```
echo -n 'sammy:' >> .htpasswd
```

Next, add an encrypted password entry for the username by typing:

```
openssl passwd -apr1 >> .htpasswd
```

You can repeat this process for additional usernames. You can see how the usernames and encrypted passwords are stored within the file by typing:

```
cat .htpasswd
```

User: sammy
Password: sammy

Output
sammy:$apr1$wI1/T0nB$jEKuTJHkTOOWkopnXqC1d1
