<p align="center">
<h1 align="center"><b>Find ClosestCommonManager</b></h1>
<p align="center">GO programming challange to find the closest common manager in organisation hierarchy</p>

## About me:

  * Abdul Ahad (Software Developer)

##  Assumptions
<p align="left">The task in completed based on following assumptions on Employee and Organisation</p>

  * Two or more Employess with same name should not be added to the organisation 
  * Each employee will have only one immediate manager. While assigning multiple Manager to same employee an error message('Manager already exist') will appear.
  * No employee will leave the organisation after joining


## Usage

#### Running the executable/binary

    user@host:~$ cd src/
    user@host:~$ go run employee.go organisation.go main.go

#### Running unit tests 

    user@host:~$ cd src/
    user@host:~$ go test -v

#### Running tests in docker containers
    user@host:~$ cd scripts

##### To build the image and run the container
    user@host:~$ sh run.sh


##### To stop the container and remove image:
    user@host:~$ sh clean.sh


#### Test cases:
1. TestEmployeeStruct
1. TestOrganisationStruct
1. TestEmployeeReportee
1. TestAddReporteeFunc
1. TestAddReporteeError
1. TestManagerExistsForEmp
1. TestAddEmployeeFunc
1. TestEmployeeIdGeneration
1. TestIsUnderFunc
1. TestClosestCommonManagerFunc
1. TestDetachedCeo
1. TestNonExistingEmployee
1. TestNilEmployees
1. TestOnlyEmployeeCeo


