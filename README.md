# 02-601CodingProject

USEFUL LINKS:
Planning Document Full: https://docs.google.com/document/d/1enU10H2dsIX9Efk3qIqXVTmk3hnsk2ALxBhD7Ja-hAE/edit?usp=sharing
Sampling Cells Planning Doc: https://docs.google.com/document/d/1cIrWe1JmXmto3qtKzb1PK6JLwCKu-wJ-fY96jvfNb1c/edit?usp=sharing
GIT BOOK: https://git-scm.com/book/en/v2

Team:
Saideep 
Kwanho 
Zhenyu
Xinling

This repository will serve as the shared repository for code related to our Programming-for-Scientists final coding project. Make sure to download git if you are new to using it/don't have it downloaded already. To get started, here are a few command line commands for collaboration using git:

1.) git clone https://github.com/SaideepGona/02-601CodingProject.git

This will download the contents of this repository into your current directory. Since I assume we will use golang, anywhere in the \src directory should be fine. It will create a new folder called 02-601CodingProject. Everything within this folder counts as part of the git repository. You are free to make any changes to your local copy of this folder, all sub-contents are considered part of the repository.

You should also run the below commands, replacing the info with your own.

  git config --global user.email "you@example.com"
  git config --global user.name "Your Name"



2.) git status

Once you have made some changes, executing this command anywhere in the repository will summarize the changes you have made since a previous commit. Commits will be covered later.

3.) git add *

This will add the changes listed in git status to a temporary "staging area". * means "all" here. Specific files can also be added individually. 

4.) git commit -m "Message"

Creates an official "save" point with a hopefully useful message. Only the changes staged with -git add- are included. Do this after making decent progress

5.) git pull

This "pulls" changes from the shared repository to your local repository. Allows you to have a more updated version. If there are differences in the files between what you have locally and what is stored remotely, git will ask you to "resolve merge conflicts". This may be confusing, so ask for help before resolving these. If you are commiting regularly, you can always restore to a previous commit if you accidently mess something up. 

6.) git push

Sends the changes up until your last commit to the shared repository on GitHub. You have to have pulled all changes and resolved conflicts with the shared branch before git will let you push your changes to it.


You don't have to use all this functionality, but in general this will make collaboration easier.


