# Scheduled Cow

Playing around with the idea of a container for scheduling jobs on kubernetes.

This project is probably going to die before its finished. Luckily, at
the time of writing, there is a proposal for [ScheduledJob](https://github.com/kubernetes/kubernetes/issues/2156)'s
waiting to be accepted in the Kubernetes repository on github.


## Production ready?

Not at all. This is a proof of concept and project aims to use and learn how
to control kubernetes using ``kubectl``. I think it would be more robust and
flexible to use something that uses the API instead. But then again, I am
not expert.


## Why Scheduled Cow

``Ku`` in Kubernetes means Cow in Norwegian. Not very cleaver and funny but at
least I tried.  


## Dependencies

All dependencies are listed in ``scripts/deps.sh`` and is manually maintained.
The format of ``deps.sh`` should be of the following format, where SHA and date
is manually at the time of download.

```sh
# commit <SHA> (<date>)
go get <repository>
```

Should probably add these vendor specific dependencies to the repository or
use ``godeps`` at some point in the future.

## Other scripts

```sh
# Download deps and runs the program
$ scripts/run.sh
```


## Todo

  * [x] Does kubectl exists on the machine
  * [ ] Run a one-off container in kube
  * [ ] Create a scheduler that start jobs every X seconds
  * [ ] Create a scheduler based on cron for startin a job at time T
  * [ ] Figure out how to specify new jobs. Redploy this container or dynamically post new jobs?
  * [ ] Figure out how to transfer data out of the cluster
