# betwixt
 A tool for testing configuration management against remote resources ad-hoc

# Portability
Currently betwixt is not portable and targets linux usage only. I've no immediate plans to change that however it should work in WSL on windows.

# Why
During the development of Ansible roles I often used the workflow of having a Vagrantfile that launched a simple t2.micro instance to run the ansible against for testing. Paired with a Makefile and simply playbook, it was wrapped up as a simple `make test` which:

1. launched instance
2. ran ansible against it
3. applied goss tests to see if ansible did what we wanted

This allowed me to run the ansible frequently during development and suss out if it was working correctly or test templating and other things rather quickly. Our testing had to be within the context of our infrastructure and hierarchial AMI structure. It also could not be reflective of our local system, and thus had to be remote. 

Vagrant, however, ultimately has a rather broad use case and required a lot of duplication and complexity to support this simple usecase. The vagrant-aws plugin is also no longer maintained and eventually will hit a horizon where it breaks without modification.

Rather than take on the burden of maintaining vagrant-aws, I opted to try and write my own tool specific to the usecase of testing automation/config management in a contained space. 

One might ask "why not terraform or docker?" and the answer is simply that sometimes you need to test within the context of a specific AMI, network, with aws or other cloud providers metadata APIs etc. It's a different use case than the one you are suggesting. 