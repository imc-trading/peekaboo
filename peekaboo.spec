%define name %NAME%
%define version %VERSION%
%define release %RELEASE%
%define buildroot %{_topdir}/BUILDROOT
%define sources %{_topdir}/SOURCES

BuildRoot: %{buildroot}
Source: %SOURCE%
Summary: %{name}
Name: %{name}
Version: %{version}
Release: %{release}
License: Apache License, Version 2.0
Group: System
AutoReqProv: no

%description
%{name}

%prep
mkdir -p %{buildroot}/usr/bin
cp %{sources}/bin/* %{buildroot}/usr/bin
mkdir -p %{buildroot}/var/lib/%{name}
cp -r %{sources}/static %{buildroot}/var/lib/%{name}
mkdir -p %{buildroot}/etc/systemd/system
cp %{sources}/files/%{name}.service %{buildroot}/etc/systemd/system/%{name}.service
mkdir -p %{buildroot}/etc/init.d
cp %{sources}/files/%{name}.initd %{buildroot}/etc/init.d/%{name}
mkdir -p %{buildroot}/etc/sysconfig
cp %{sources}/files/%{name} %{buildroot}/etc/sysconfig/%{name}

%post
which systemctl &>/dev/null && systemctl daemon-reload

%preun
# Disable and stop on uninstall
if [ "${1}" == "0" ]; then
  if which systemctl &>/dev/null; then
    systemctl stop peekaboo
    systemctl disable peekaboo
  else
    service peekaboo stop
    chkconfig program off
  fi
fi

%postun
# Restart on upgrade
if [ "${1}" == "1" ]; then
  if which systemctl &>/dev/null; then
    systemctl restart peekaboo
  else
    service peekaboo restart
  fi
fi

%files
%defattr(-,root,root)
/usr/bin/%{name}
/var/lib/%{name}
/etc/systemd/system/%{name}.service
%(755,-,-) /etc/init.d/%{name}
%config(noreplace) /etc/sysconfig/%{name}
