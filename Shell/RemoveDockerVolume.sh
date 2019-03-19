docker volume rm $(docker volume ls | egrep -o " [^ Nm]*$")
