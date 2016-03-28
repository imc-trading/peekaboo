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

%files
%defattr(-,root,root)
/usr/bin/%{name}
/var/lib/%{name}
/etc/systemd/system/%{name}.service
%(755,-,-) /etc/init.d/%{name}
%config(noreplace) /etc/sysconfig/%{name}
