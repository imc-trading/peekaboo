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
cp -r %{sources}/templates %{buildroot}/var/lib/%{name}
cp -r %{sources}/public %{buildroot}/var/lib/%{name}

%files
%defattr(-,root,root)
/usr/bin
/var/lib/%{name}
