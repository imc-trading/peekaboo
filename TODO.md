# General
- Support for plugins i.e. arbitrary endpoints
- Endpoint for enabled plugins "/plugins"
- Configuration file for endpoints/ timeouts, hide/ show
- Improve documentation

# Security
- TLS 1.2+
- Split daemon in root/ non-root using a socket for communication or similar solution to minimize security risks

# Backend
- Improve IPMI for diff. HW
- Keep history for metrics allow for Dashboards file per day or circular db?
- Nw Interface cable S/N
- Add endpoint for disk space "df"

# Message Bus
- Re-add Kafka support
- Kafka events on changes use JSON PATCH

# Front-end
- Button for horizontal table in Front-end
- Front-end check if a plugin is enabled/ available otherwise hide the page

# Minor
- Don't close Dropdown for columns on click
- REST call for Set Timeout
