let cacheData = "appV1";


this.addEventListener("install", (event) =>{
    event.waitUntil(
        caches.open(cacheData).then((cache) =>{
            cache.addAll([
                '/static/js/main.dc75e769.js',
                '/static/css/main.4107fc36.css',
                '/static/media/DMSans-Regular.067ebd7ed3c947d82dc4.ttf',
                '/favicon.ico',
                '/apple-touch-icon.png',
                '/logo-white.png',
                '/manifest.json',
                '/index.html',
                '/automobile_dealer'
            ])
        })
    )
})


this.addEventListener("fetch", (event) =>{
   if(!navigator.onLine){
    event.respondWith(
        caches.match(event.request).then((resp) =>{
            if(resp){
                return resp
            }

            let requestUrl = event.request.clone()
            fetch(requestUrl)
        })
    )
   }
})