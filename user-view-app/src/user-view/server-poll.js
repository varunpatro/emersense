(function(global) {

    function normalize(url) {
        var strippedURL = new URL(url);
        strippedURL.search = '';
        return strippedURL.toString();
    }

    global.pollServer = function(request) {
        // Attempt to fetch request.
        return global.fetch(request).then(function(res) {
            if (res.ok) {
                return global.caches.open(global.toolbox.options.cacheName).then(function(cache) {
                    return cache.put(normalize(request.url), res.clone()).then(function(){
                        return res;
                    });
                });
            }
            throw new Error("A response with error status code was returned");
        }).catch(function(error) {
            return global.caches.open(global.toolbox.options.cacheName).then(function(cache) {
                return cache.match(normalize(request.url));
            }); 
        });
    }
   
})(self);
