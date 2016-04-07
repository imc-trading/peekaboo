/*
 * Application
 */

var app = angular.module('peekaboo', [
  'ngRoute',
  'ngResource',
  'ui.bootstrap',
  'smart-table',
  'ngFlash',
  'ngAnimate'
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

app.controller('PageCtrl', ['Flash', function(Flash) {
  console.log("Page Controller reporting for duty.");
  Flash.clear();
}]);

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
  return function (input,restr,newstr) {
    var re = new RegExp(restr);
    return input.replace(re, newstr)
  };
});

/*
 * Controllers
 */

// CPU
app.controller('cpuController', [ '$scope', '$resource', 'Flash',  function($scope, $resource, Flash) {
  var resource = $resource('/api/system/cpu');

  resource.get().$promise.then(function(value) {
    $scope.cpu = value;
//    console.log (value); 
  }, function(err) {
    var msg = "<strong>Failed to request URL</strong>: " + err.config.url + " <strong>error</strong>: " + err.data;
    var id = Flash.create('danger', msg, 10000, {class: 'custom-class', id: 'custom-id'}, true);
  });
} ]);

// Memory
app.controller('memoryController', [ '$scope', '$resource', 'Flash', function($scope, $resource, Flash) {
  var resource = $resource('/api/system/memory');

  resource.get().$promise.then(function(value) {
    $scope.memory = value;
//    console.log (value); 
  }, function(err) {
    var msg = "<strong>Failed to request URL</strong>: " + err.config.url + " <strong>error</strong>: " + err.data;
    var id = Flash.create('danger', msg, 10000, {class: 'custom-class', id: 'custom-id'}, true);
  });
} ]);

// System
app.controller('systemController', [ '$scope', '$resource', 'Flash', function($scope, $resource, Flash) {
  var resource = $resource('/api/system');

  resource.get().$promise.then(function(value) {
    $scope.system = value;
//    console.log (value); 
  }, function(err) {
    var msg = "<strong>Failed to request URL</strong>: " + err.config.url + " <strong>error</strong>: " + err.data;
    var id = Flash.create('danger', msg, 10000, {class: 'custom-class', id: 'custom-id'}, true);
  });
} ]);

// OS
app.controller('osController', [ '$scope', '$resource', 'Flash', function($scope, $resource, Flash) {
  var resource = $resource('/api/system/os');

  resource.get().$promise.then(function(value) {
    $scope.os = value;
//    console.log (value); 
  }, function(err) {
    var msg = "<strong>Failed to request URL</strong>: " + err.config.url + " <strong>error</strong>: " + err.data;
    var id = Flash.create('danger', msg, 10000, {class: 'custom-class', id: 'custom-id'}, true);
  });
} ]);

// Interfaces
app.controller('interfacesController', [ '$scope', '$resource', 'Flash', function($scope, $resource, Flash) {
  var resource = $resource('/api/network/interfaces');

  resource.query().$promise.then(function(value) {
    $scope.interfaces = value;
//    console.log (value); 
  }, function(err) {
    var msg = "<strong>Failed to request URL</strong>: " + err.config.url + " <strong>error</strong>: " + err.data;
    var id = Flash.create('danger', msg, 10000, {class: 'custom-class', id: 'custom-id'}, true);
  });
} ]);

// Routes
app.controller('routesController', [ '$scope', '$resource', 'Flash', function($scope, $resource, Flash) {
  var resource = $resource('/api/network/routes');

  resource.query().$promise.then(function(value) {
    $scope.routes = value;
//    console.log (value);
  }, function(err) {
    var msg = "<strong>Failed to request URL</strong>: " + err.config.url + " <strong>error</strong>: " + err.data;
    var id = Flash.create('danger', msg, 10000, {class: 'custom-class', id: 'custom-id'}, true);
  });
} ]);

// Disks
app.controller('disksController', [ '$scope', '$resource', 'Flash', function($scope, $resource, Flash) {
  var resource = $resource('/api/storage/disks');

  resource.query().$promise.then(function(value) {
    $scope.disks = value;
//    console.log (value);
  }, function(err) {
    var msg = "<strong>Failed to request URL</strong>: " + err.config.url + " <strong>error</strong>: " + err.data;
    var id = Flash.create('danger', msg, 10000, {class: 'custom-class', id: 'custom-id'}, true);
  });
} ]);

// Mounts
app.controller('mountsController', [ '$scope', '$resource', 'Flash', function($scope, $resource, Flash) {
  var resource = $resource('/api/storage/mounts');

  resource.query().$promise.then(function(value) {
    $scope.mounts = value;
//    console.log (value);
  }, function(err) {
    var msg = "<strong>Failed to request URL</strong>: " + err.config.url + " <strong>error</strong>: " + err.data;
    var id = Flash.create('danger', msg, 10000, {class: 'custom-class', id: 'custom-id'}, true);
  });
} ]);

// LVM Physical Volumes
app.controller('physVolsController', [ '$scope', '$resource', 'Flash', function($scope, $resource, Flash) {
  var resource = $resource('/api/storage/lvm/physvols');

  resource.query().$promise.then(function(value) {
    $scope.physVols = value;
//    console.log (value);
  }, function(err) {
    var msg = "<strong>Failed to request URL</strong>: " + err.config.url + " <strong>error</strong>: " + err.data;
    var id = Flash.create('danger', msg, 10000, {class: 'custom-class', id: 'custom-id'}, true);
  });
} ]);

// LVM Logical Volumes
app.controller('logVolsController', [ '$scope', '$resource', 'Flash', function($scope, $resource, Flash) {
  var resource = $resource('/api/storage/lvm/logvols');

  resource.query().$promise.then(function(value) {
    $scope.logVols = value;
//    console.log (value);
  }, function(err) {
    var msg = "<strong>Failed to request URL</strong>: " + err.config.url + " <strong>error</strong>: " + err.data;
    var id = Flash.create('danger', msg, 10000, {class: 'custom-class', id: 'custom-id'}, true);
  });
} ]);

// LVM Volume Groups
app.controller('volGrpsController', [ '$scope', '$resource', 'Flash', function($scope, $resource, Flash) {
  var resource = $resource('/api/storage/lvm/volgrps');

  resource.query().$promise.then(function(value) {
    $scope.volGrps = value;
//    console.log (value);
  }, function(err) {
    var msg = "<strong>Failed to request URL</strong>: " + err.config.url + " <strong>error</strong>: " + err.data;
    var id = Flash.create('danger', msg, 10000, {class: 'custom-class', id: 'custom-id'}, true);
  });
} ]);

// Sysctls
app.controller('sysctlsController', [ '$scope', '$resource', 'Flash', function($scope, $resource, Flash) {
  var resource = $resource('/api/system/sysctls');

  resource.query().$promise.then(function(value) {
    $scope.sysctls = value;
//    console.log (value);
  }, function(err) {
    var msg = "<strong>Failed to request URL</strong>: " + err.config.url + " <strong>error</strong>: " + err.data;
    var id = Flash.create('danger', msg, 10000, {class: 'custom-class', id: 'custom-id'}, true);
  });
} ]);

// Docker
app.controller('dockerController', [ '$scope', '$resource', 'Flash', function($scope, $resource, Flash) {
  var resource = $resource('/api/docker');

  resource.get().$promise.then(function(value) {
    $scope.docker = value;
//    console.log (value);
  }, function(err) {
    var msg = "<strong>Failed to request URL</strong>: " + err.config.url + " <strong>error</strong>: " + err.data;
    var id = Flash.create('danger', msg, 10000, {class: 'custom-class', id: 'custom-id'}, true);
  });
} ]);

// Images
app.controller('imagesController', [ '$scope', '$resource', 'Flash', function($scope, $resource, Flash) {
  var resource = $resource('/api/docker/images');

  resource.query().$promise.then(function(value) {
    $scope.images = value;
//    console.log (value);
  }, function(err) {
    var msg = "<strong>Failed to request URL</strong>: " + err.config.url + " <strong>error</strong>: " + err.data;
    var id = Flash.create('danger', msg, 10000, {class: 'custom-class', id: 'custom-id'}, true);
  });
} ]);

// Containers
app.controller('containersController', [ '$scope', '$resource', 'Flash', function($scope, $resource, Flash) {
  var resource = $resource('/api/docker/containers');

  resource.query().$promise.then(function(value) {
    $scope.containers = value;
//    console.log (value);
  }, function(err) {
    var msg = "<strong>Failed to request URL</strong>: " + err.config.url + " <strong>error</strong>: " + err.data;
    var id = Flash.create('danger', msg, 10000, {class: 'custom-class', id: 'custom-id'}, true);
  });
} ]);

// Containers
app.controller('ipmiController', [ '$scope', '$resource', 'Flash', function($scope, $resource, Flash) {
  var resource = $resource('/api/system/ipmi');

  resource.get().$promise.then(function(value) {
    $scope.ipmi = value;
//    console.log (value);
  }, function(err) {
    var msg = "<strong>Failed to request URL</strong>: " + err.config.url + " <strong>error</strong>: " + err.data;
    var id = Flash.create('danger', msg, 10000, {class: 'custom-class', id: 'custom-id'}, true);
  });
} ]);
