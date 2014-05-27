## tmate tmux sharing buildpack

This buildpack will install [tmate](http://tmate.io) and then launch it, revealing
the url to connect to via ssh to gain access to the application container.

##tl;dr

- Create a folder, with a manifest.yml and an arbitrary file (something for cf to upload)
- Edit the manifest;


```yaml
---
applications:
- name: tmate-test
  memory: 256M
  instances: 1
  host: tmate-test
  domain: cfapps.io
  path: .
  buildpack: https://github.com/danhigham/cloudfoundry-tmate-buildpack.git
  command: launch
```

- Push it (cf push)

- Retrieve the ssh url for your tmate session;

```
$ cf logs tmate-test --recent
....
....

2014-05-27T14:32:37.93-0700 [STG]     OUT -----> Uploading droplet (6.0M)
2014-05-27T14:32:43.26-0700 [DEA]     OUT Starting app instance (index 0) with guid c94f71bc-7077-4dca-a578-11cd5b9200d0
2014-05-27T14:32:44.88-0700 [App/0]   ERR 2014/05/27 21:32:44 Starting tmate...
2014-05-27T14:32:44.88-0700 [App/0]   ERR 2014/05/27 21:32:44 1000
2014-05-27T14:32:44.88-0700 [App/0]   ERR 2014/05/27 21:32:44 1000
2014-05-27T14:32:48.12-0700 [App/0]   ERR 2014/05/27 21:32:48 gd************DdV@ny.tmate.io

```

- ssh to the tmate address

```
$ ssh gd************DdV@ny.tmate.io
```

Even better, hand the URL out to your friends and collaborate on a tmate session!

Be aware that when a participant exits the shell, the session is done. However,
to start it again, just restart the app!
