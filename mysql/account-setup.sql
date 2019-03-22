/* account for web */;
CREATE USER docker1 IDENTIFIED WITH mysql_native_password BY 'docker1';
GRANT ALL PRIVILEGES ON crr.* TO docker1;

/* account for use when you want to modify db content_type_id
   note that custom content will be lost if you remove docker volumin */;
CREATE USER 'contentadmin'@'%' IDENTIFIED WITH mysql_native_password BY 'contentadmin';
GRANT ALL PRIVILEGES ON crr.* TO contentadmin;
