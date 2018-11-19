@echo off

set "CATALINA_HOME=%CD%\server\apache-tomcat-8.0.47"

echo "starting camunda BPM platform 7.8.0 on Apache Tomcat 8.0.47"

cd server\apache-tomcat-8.0.47\bin\
start startup.bat

ping -n 5 localhost > NULL
start http://localhost:8080/camunda-welcome/index.html
 