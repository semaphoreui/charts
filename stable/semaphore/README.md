# Semaphore UI official chart

![Version: 16.0.0](https://img.shields.io/badge/Version-16.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2.15.0](https://img.shields.io/badge/AppVersion-2.15.0-informational?style=flat-square)

Modern UI and powerful API for Ansible, Terraform, OpenTofu, PowerShell and other DevOps tools

**Homepage:** <https://semaphoreui.com/>

## Installing the Chart

```console
helm repo add semaphoreui https://semaphoreui.github.io/charts
helm repo update
```

**Open-source version:**

```console
helm install semaphore semaphoreui/semaphore
```

**Pro version:**

```console
helm install semaphore semaphoreui/semaphore --set image.repository=public.ecr.aws/semaphore/pro/server
```

## Installing OCI from GitHub Container Registry (GitHub authentication required)

**Open-source version:**

```console
helm install semaphore oci://ghcr.io/semaphoreui/charts/semaphore
```

**Pro version:**

```console
helm install semaphore oci://ghcr.io/semaphoreui/charts/semaphore --set image.repository=public.ecr.aws/semaphore/pro/server
```

## Example for Values

### Create `admin` user

```yml
admin:
  create: true
  password: changeme
```

Now you can login with username `admin` and password `changeme`.

### Ingress Enabled

```yml
ingress:
  enabled: false

  hosts:
    - host: semaphore.example.com
      paths:
        - path: /
          pathType: Prefix
```

### Bundled MariaDB

```console
database:
  type: mysql

  host: mariadb
  port: 3306

  password: p455w0rd

mariadb:
  enabled: true

  auth:
    password: p455w0rd
```

### Bundled PostgreSQL

```console
database:
  type: postgres

  host: postgresql
  port: 5432

  usernameFromSecret: false
  passwordKey: password
  existingSecret: postgresql

postgresql:
  enabled: true

  auth:
    password: p455w0rd
```

### OpenID Connect

```console
oidc:
  enable: true

  providers:
    keycloak:
      display_name: Keycloak
      provider_url: https://auth.example.com/auth/realms/example
      redirect_url: https://semaphore.example.com/api/auth/oidc/keycloak/redirect
      client_id: semaphore
      client_secret: 0208901c-ecd7-46ae-931a-d03f02e8dcd2
      username_claim: preferred_username
      name_claim: preferred_username
      email_claim: email
```

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| Denis Gukov | <denis@semaphoreui.com> | <https://github.com/fiftin> |

## Source Code

* <https://github.com/semaphoreui/semaphore>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| oci://registry-1.docker.io/bitnamicharts | mariadb | 20.4.2 |
| oci://registry-1.docker.io/bitnamicharts | postgresql | 16.6.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| admin.create | bool | `false` | Create an local admin user |
| admin.email | string | `"admin@localhost"` | Email for local admin |
| admin.emailKey | string | `"email"` | Key used within secret for email |
| admin.existingSecret | string | `nil` | Existing secret to use for admin |
| admin.fullname | string | `"Admin"` | Fullname for local admin |
| admin.fullnameKey | string | `"fullname"` | Key used within secret for fullname |
| admin.password | string | `nil` | Password for local admin |
| admin.passwordKey | string | `"password"` | Key used within secret for password |
| admin.username | string | `"admin"` | Username for local admin |
| admin.usernameKey | string | `"username"` | Key used within secret for username |
| affinity | object | `{}` | Affinity for the deployment |
| annotations | object | `{}` | Define additional annotations |
| config.forwarded_env_vars | list | `[]` | List of environment variables to forward into the application environment. Modify this list if additional environment variables need to be accessible. |
| database.existingSecret | string | `nil` | Existing secret to use for credentials |
| database.host | string | `nil` | Host for database connection |
| database.name | string | `"semaphore"` | Name of the used database |
| database.options | object | `{}` | Options for database connection |
| database.password | string | `nil` | Password for database |
| database.passwordKey | string | `"password"` | Key used within secret for password |
| database.path | string | `"/var/lib/semaphore/database.boltdb"` | Path for the boltdb |
| database.persistence.accessModes | list | `["ReadWriteOnce"]` | Access modes used for boltdb volume |
| database.persistence.enabled | bool | `true` | Enable persistence for boltdb |
| database.persistence.existingClaim | string | `nil` | Name of an already existing claim |
| database.persistence.size | string | `"5G"` | Size for boltdb volume |
| database.persistence.storageClass | string | `nil` | Storage class used for boltdb volume |
| database.port | string | `nil` | Port for database connection |
| database.type | string | `"bolt"` | Type of database backend |
| database.username | string | `"semaphore"` | Username for database |
| database.usernameFromSecret | bool | `true` | Read username from secret |
| database.usernameKey | string | `"username"` | Key used within secret for username |
| email.alert | bool | `false` | Enable email alerting |
| email.existingSecret | string | `nil` | Existing secret to use for email |
| email.host | string | `nil` | Host of the SMTP server |
| email.password | string | `nil` | Password for SMTP server |
| email.passwordKey | string | `"password"` | Key used within secret for password |
| email.port | string | `nil` | Port of the SMTP server |
| email.secure | bool | `false` | Enable a secure connection |
| email.sender | string | `nil` | Sender for email alerting |
| email.username | string | `nil` | Username for SMTP server |
| email.usernameKey | string | `"username"` | Key used within secret for username |
| envFromConfigMaps | list | `[]` | List of environment variables from existing configmaps |
| envFromSecrets | list | `[]` | List of environment variables from existing secrets |
| extraEnvSecrets | object | `{}` | Extra environment variables from secrets |
| extraEnvVariables | object | `{}` | Extra environment variables from mapping |
| extraInitContainers | list | `[]` | List of extra init containers |
| extraSidecarContainers | list | `[]` | List of extra sidecar containers |
| extraVolumeMounts | list | `[]` | List of extra volume mounts |
| extraVolumes | list | `[]` | List of extra volumes |
| fullnameOverride | string | `""` | Override the fullname |
| general.additionalPythonPackages | list | `[]` | Additional Python packages |
| general.gitClient | string | `"cmd_git"` | Use Git client implementation |
| general.host | string | `nil` | Host to access Semaphore |
| general.maxParallelTasks | int | `0` | Maximum parallel tasks |
| general.nonAdminCanCreateProject | bool | `false` | Allow non-admins to create projects |
| general.passwordLoginDisable | bool | `false` | Disable password login |
| general.sshConfigPath | string | `nil` | Path to SSH config |
| general.tmpPath | string | `"/tmp/semaphore"` | Working directory for Semaphore |
| general.useRemoteRunner | bool | `false` | Enable usage of remote runners |
| image.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| image.pullSecrets | list | `[]` | Optional name of pull secret if using a private registry |
| image.repository | string | `"semaphoreui/semaphore"` | Image repository used by deployment |
| image.tag | string | `""` | Optional tag for the repository, defaults to app version |
| ingress.annotations | object | `{}` | Additional annotations for the ingress |
| ingress.className | string | `nil` | Class name for the ingress resource |
| ingress.enabled | bool | `false` | Enable ingress |
| ingress.hosts | list | `[{"host":"example.local","paths":[{"path":"/","pathType":"Prefix"}]}]` | Host definition for ingress |
| ingress.labels | object | `{}` | Additional labels for the ingress |
| ingress.tls | list | `[]` | Optional TLS configuration for ingress |
| labels | object | `{}` | Define additional labels |
| ldap.binddn | string | `nil` | BindDN for LDAP authentication |
| ldap.binddnKey | string | `"username"` | Key used within secret for username |
| ldap.enable | bool | `false` | Enable LDAP authentication |
| ldap.existingSecret | string | `nil` | Existing secret to use for ldap |
| ldap.filter | string | `nil` | Search filter for LDAP |
| ldap.mappings | object | `{"cn":"cn","dn":"dn","mail":"mail","uid":"uid"}` | Mapping for LDAP attributes |
| ldap.needtls | bool | `false` | Enable TLS connection to LDAP |
| ldap.password | string | `nil` | Password for LDAP authentication |
| ldap.passwordKey | string | `"password"` | Key used within secret for username |
| ldap.searchdn | string | `nil` |  |
| ldap.server | string | `nil` | Address of LDAP server |
| mariadb.architecture | string | `"standalone"` | Architecture for mariadb |
| mariadb.auth.database | string | `"semaphore"` | Database created for semaphore |
| mariadb.auth.password | string | `nil` | Password for semaphore database |
| mariadb.auth.username | string | `"semaphore"` | Username for semaphore database |
| mariadb.enabled | bool | `false` | Enable mariadb dependency |
| mariadb.fullnameOverride | string | `"mariadb"` | Override fullname of mariadb dependency |
| mariadb.metrics.enabled | bool | `true` | Enable metrics for mariadb |
| mariadb.metrics.serviceMonitor.enabled | bool | `false` | Enable service monitor for mariadb |
| mariadb.serviceAccount.create | bool | `true` | Create service account for mariadb |
| nameOverride | string | `""` | Override the name |
| nodeSelector | object | `{}` | Node selector for the deployment |
| oidc.enable | bool | `false` | Enable oidc authentication |
| oidc.providers | object | `{}` | Dictionary of oidc providers |
| persistence.accessModes | list | `["ReadWriteOnce"]` | Access modes used for workdir volume |
| persistence.enabled | bool | `true` | Enable persistence for workdir |
| persistence.existingClaim | string | `nil` | Name of an already existing claim |
| persistence.size | string | `"10G"` | Size for boltdb volume |
| persistence.storageClass | string | `nil` | Storage class used for workdir volume |
| podSecurityContext | object | `{}` | Security context for the pod |
| postgresql.architecture | string | `"standalone"` | Architecture for postgresql |
| postgresql.auth.database | string | `"semaphore"` | Database created for semaphore |
| postgresql.auth.password | string | `nil` | Password for semaphore database |
| postgresql.auth.username | string | `"semaphore"` | Username for semaphore database |
| postgresql.enabled | bool | `false` | Enable postgresql dependency |
| postgresql.fullnameOverride | string | `"postgresql"` | Override fullname of postgresql dependency |
| postgresql.metrics.enabled | bool | `true` | Enable metrics for postgresql |
| postgresql.metrics.serviceMonitor.enabled | bool | `false` | Enable service monitor for postgresql |
| postgresql.serviceAccount.create | bool | `true` | Create service account for postgresql |
| replicaCount | int | `1` | Replicas for the deployment |
| resources | object | `{"limits":{},"requests":{"cpu":"100m","memory":"64Mi"}}` | Resources for the deployment |
| runner.existingSecret | string | `nil` | Existing secret to use for runner |
| runner.token | string | `nil` | Runner registration token |
| runner.tokenKey | string | `"token"` | Key used within secret for token |
| secrets.accesskeyEncryption | string | `nil` | Access key encryption secret, generated if not present |
| secrets.accesskeyEncryptionKey | string | `"accesskeyEncryption"` | Key used within secret for accesskeyEncryption |
| secrets.cookieEncryption | string | `nil` | Cookie encryption secret, generated if not present |
| secrets.cookieEncryptionKey | string | `"cookieEncryption"` | Key used within secret for cookieEncryption |
| secrets.cookieHash | string | `nil` | Cookie hash secret, generated if not present |
| secrets.cookieHashKey | string | `"cookieHash"` | Key used within secret for cookieHash |
| secrets.existingSecret | string | `nil` | Existing secret to use for secrets |
| securityContext | object | `{"fsGroup":1001}` | Security context for the deployment |
| service.annotations | object | `{}` | Additional annotations for the service |
| service.internalPort | int | `3000` | Internal port of the service |
| service.labels | object | `{}` | Additional labels for the service |
| service.port | int | `3000` | Port of the service |
| service.type | string | `"ClusterIP"` | Type of the service |
| serviceAccount.annotations | object | `{}` | Define annotations for the service account |
| serviceAccount.create | bool | `true` | Create a new service account |
| serviceAccount.name | string | `""` | Optional name for an existing service account |
| slack.alert | bool | `false` | Enable slack alerting |
| slack.existingSecret | string | `nil` | Existing secret to use for slack |
| slack.url | string | `nil` | URL used for slack |
| slack.urlKey | string | `"url"` | Key used within secret for url |
| telegram.alert | bool | `false` | Enable telegram alerting |
| telegram.chat | string | `nil` | Chat used for telegram |
| telegram.chatKey | string | `"chat"` | Key used within secret for chat |
| telegram.existingSecret | string | `nil` | Existing secret to use for telegram |
| telegram.token | string | `nil` | Token used for telegram |
| telegram.tokenKey | string | `"token"` | Key used within secret for token |
| tolerations | list | `[]` | Tolerations for the deployment |
| updateStrategy | object | `{"type":"Recreate"}` | Updaqte strategy for deployment |
