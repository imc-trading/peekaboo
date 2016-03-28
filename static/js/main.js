var app = angular.module('peekaboo', [
  'ngRoute',
  'ngResource',
  'ui.bootstrap',
  'smart-table'
])

/*
 * Routes
 */

app.config(['$routeProvider', function ($routeProvider) {
  $routeProvider
    .when("/", {templateUrl: "partials/dashboard.html", controller: "PageCtrl"})
    .when("/network", {templateUrl: "partials/network.html", controller: "PageCtrl"})
    .when("/storage", {templateUrl: "partials/storage.html", controller: "PageCtrl"})
    .when("/system", {templateUrl: "partials/system.html", controller: "PageCtrl"})
    .when("/docker", {templateUrl: "partials/docker.html", controller: "PageCtrl"})
}]);

/**
 * Controls all other Pages
 */
app.controller('PageCtrl', function (/* $scope, $location, $http */) {
  console.log("Page Controller reporting for duty.");
});

/*
 * Filters
 */

// joinBy
app.filter('joinBy', function () {
  return function (input,delimiter) {
    return (input || []).join(delimiter || ',');
  };
});

// replace
app.filter('replace', function () {
  return function (input,oldstr,newstr) {
    var re = new RegExp(oldstr, "g");
    return input.replace(re, newstr)
  };
});

/*
 * Controllers
 */

// CPU
app.controller('cpuController', [ '$scope', '$resource', function($scope, $resource) {
  var resource = $resource('/api/v1/system/cpu');

  resource.get().$promise.then(function(value) {
    $scope.cpu = value;
//    console.log (value); 
  });
} ]);

// Memory
app.controller('memoryController', [ '$scope', '$resource', function($scope, $resource) {
  var resource = $resource('/api/v1/system/memory');

  resource.get().$promise.then(function(value) {
    $scope.memory = value;
//    console.log (value); 
  });
} ]);

// System
app.controller('systemController', [ '$scope', '$resource', function($scope, $resource) {
  var resource = $resource('/api/v1/system');

  resource.get().$promise.then(function(value) {
    $scope.system = value;
//    console.log (value); 
  });
} ]);

// OS
app.controller('osController', [ '$scope', '$resource', function($scope, $resource) {
  var resource = $resource('/api/v1/system/os');

  resource.get().$promise.then(function(value) {
    $scope.os = value;
//    console.log (value); 
  });
} ]);

// Interfaces
app.controller('interfacesController', [ '$scope', '$resource', function($scope, $resource) {
  var resource = $resource('/api/v1/network/interfaces');

  resource.query().$promise.then(function(value) {
    $scope.interfaces = value;
//    console.log (value); 
  });
} ]);

// Routes
app.controller('routesController', [ '$scope', '$resource', function($scope, $resource) {
  var resource = $resource('/api/v1/network/routes');

  resource.query().$promise.then(function(value) {
    $scope.routes = value;
//    console.log (value);
  });
} ]);

// Disks
app.controller('disksController', [ '$scope', '$resource', function($scope, $resource) {
  var resource = $resource('/api/v1/storage/disks');

  resource.query().$promise.then(function(value) {
    $scope.disks = value;
//    console.log (value);
  });
} ]);

// Mounts
app.controller('mountsController', [ '$scope', '$resource', function($scope, $resource) {
  var resource = $resource('/api/v1/storage/mounts');

  resource.query().$promise.then(function(value) {
    $scope.mounts = value;
//    console.log (value);
  });
} ]);

// LVM Physical Volumes
app.controller('physVolsController', [ '$scope', '$resource', function($scope, $resource) {
  var resource = $resource('/api/v1/storage/lvm/physvols');

  resource.query().$promise.then(function(value) {
    $scope.physVols = value;
//    console.log (value);
  });
} ]);

// LVM Logical Volumes
app.controller('logVolsController', [ '$scope', '$resource', function($scope, $resource) {
  var resource = $resource('/api/v1/storage/lvm/logvols');

  resource.query().$promise.then(function(value) {
    $scope.logVols = value;
//    console.log (value);
  });
} ]);

// LVM Volume Groups
app.controller('volGrpsController', [ '$scope', '$resource', function($scope, $resource) {
  var resource = $resource('/api/v1/storage/lvm/volgrps');

  resource.query().$promise.then(function(value) {
    $scope.volGrps = value;
//    console.log (value);
  });
} ]);

// Sysctls
app.controller('sysctlsController', [ '$scope', '$resource', function($scope, $resource) {
  var resource = $resource('/api/v1/system/sysctls');

  resource.query().$promise.then(function(value) {
    $scope.sysctls = value;
//    console.log (value);
  });
} ]);

// Docker
app.controller('dockerController', [ '$scope', '$resource', function($scope, $resource) {
  var resource = $resource('/api/v1/docker');

  resource.get().$promise.then(function(value) {
    $scope.docker = value;
//    console.log (value);
  });
} ]);

// Images
app.controller('imagesController', [ '$scope', '$resource', function($scope, $resource) {
  var resource = $resource('/api/v1/docker/images');

  resource.query().$promise.then(function(value) {
    $scope.images = value;
//    console.log (value);
  });
} ]);

// Containers
app.controller('containersController', [ '$scope', '$resource', function($scope, $resource) {
  var resource = $resource('/api/v1/docker/containers');

  resource.query().$promise.then(function(value) {
    $scope.containers = value;
//    console.log (value);
  });
} ]);
