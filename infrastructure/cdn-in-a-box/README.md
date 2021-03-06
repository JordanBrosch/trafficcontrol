<!--
    Licensed to the Apache Software Foundation (ASF) under one
    or more contributor license agreements.  See the NOTICE file
    distributed with this work for additional information
    regarding copyright ownership.  The ASF licenses this file
    to you under the Apache License, Version 2.0 (the
    "License"); you may not use this file except in compliance
    with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing,
    software distributed under the License is distributed on an
    "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
    KIND, either express or implied.  See the License for the
    specific language governing permissions and limitations
    under the License.
-->

CDN In a Box (containerized)
============================

This is intended to simplify the process of creating a "CDN in a box",  easing
the barrier to entry for newcomers as well as providing a way to spin up a
minimal CDN for full system testing.

For now,  only `traffic_ops` and `traffic_portal` are implemented.  Other
components will follow as well as details on specific parts of the
implementation.. 

To start it, install `docker-ce` and `docker-compose`.  Create an rpm for each
component and copy into the respective directory with no version number, e.g.
`traffic_ops/traffic_ops.rpm`, `traffic_portal/traffic_portal.rpm`.
    
    cd infrastructure/cdn-in-a-box
    docker-compose up --build

The process creates containers for each component with ports exposed on the host.  The
following should be available once the system is running:
   
    Traffic Portal: https://localhost    
    Traffic Ops (go): https://localhost:6443
    Traffic Ops (perl): https://localhost:60443
    Postgres: `psql -h localhost -p 5432 -U postgres`
