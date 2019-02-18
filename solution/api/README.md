### Preface

So since I'm not usually coding APIs from scratch, this was a kind of exploration mission for me. As you can see below, I have made a couple of assumptions and subjeсtive decisions:

* Response codes are a bit different from what's specified in *Swagger* definition, tried to make them more detailed for the client.

* All *Passenger* struct fields are required, so it's impossible to create a new passenger without filling out every field. Also, according to RFC, PUT should replace the whole object, so here we're also making all object fields required. In the ideal scenario the validation logging should be improved (see *Potential Improvements*)

* Decision to choose MongoDB was straight forward for me - from previous experience, it's a good choice when it comes to storing user data. I feel like its native *Objectid* field type has a lot of advantages compared to barebone *uuid*, so I've decided not to add *uuid* field (well, actually i did implement it initially by enriching database via mongo shell and .js script which generated *uuid* in newest *Binary x04* format for every document in collection, but dropped it in the end) and to use *Objectid* instead (criticism is welcomed ¯\\\_(ツ)_/¯ ).

* Mock data seeding process itself is overcomplicated a bit - I have redefined *CSV* fields manually since mongoimport can't "guess" field types precisely enough.

* I don't like the concept of serving *API* and *Web UI* on the same endpoint and making decisions based on *Content-Type* header so I would rather serve them on different endpoints (*Web UI* is not implemented so yeah, whatever, *API* could be done on */* instead of */api* )

### Potential improvements (the most obvious ones)

* [CODE] Write meaningful unit and integration tests ¯\\\_(ツ)_/¯
* [CODE] Rewrite *Update* handler so it does not look that crappy.
* [CODE] Structure and improve application configuration (env variables -> configuration file -> k\v configuration service)
* [CODE] Improve in-app logging (*JSON* format?) and feedback for the client (eg. *JSON* payload validation errors in addition to *400* response code).
* [CODE] Implement durable logic to maintain *MongoDB* session (some of mgo's calls can cause panic, so it kinda works fine out-of-the-box, but for production, usage can be slightly improved). Also, get rid of *service ordering* requirement by implementing verbose *retries* and *circuit breakers* when talking with *MongoDB*
* [CODE] Add transport layer *TLS\SSL* encryption when talking with *MongoDB*
* [CODE] To be perfectly honest and transparent - it's not like I'm doing full-stack gigs every day, so it could take another couple of hours to implement simple *Web UI* for this *API* without having a regular practice.
* [CODE] Instrument code with *Prometheus* metrics
* [CODE] Instrument code with *distributed tracing* library (*Jaeger* is my favorite at the moment)
* [CONTAINER] Enforce container security by configuring *NetworkPolicies*, *PodSecurityPolicies*, *Linux capabilities* etc.
* [CONTAINER/K8S] Test and improve *Docker\Kubernetes* *SIGTERM/SIGKILL* signals handling
* [K8S] Template *Kubernetes* manifests with *Helm\kustomize\whatever* (*kustomize* is integrated into newer *kubectl* versions, *Helm* is not just a templating engine and it has its own pros and cons)
* [K8S] Think about proper *readiness\liveness* probes - eg. implement */healthz* endpoint and decide if it's needed to include *MongoDB* connectivity check. This is a big holy war topic - some of the big companies are not using *readiness\liveness* at all intentionally - in some cases, they cause more issues than they can solve.
* [K8S] *HorizontalPodAutoscaler* based on *Prometheus* app metrics (*custom metrics* via adapter) or the ones from *Ingress* (*object metrics*)
* [K8S] Productionalize *MongoDB Kubernetes* setup (stable *Helm* chart can be used as a starting point)

* [CODE] Potentially go mad untill you achieve 101% *container and Kubernetes nativeness* with Watch method on API via WebSockets, hot config reloads e
