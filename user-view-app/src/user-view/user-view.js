Polymer({
    is: 'user-view',
    properties: {
        uuid: {
            type: String,
            value: "",
            notify: true
        },
        isSafe: {
            type: Boolean,
            value: false,
            notify: true
        },
        marked: {
            type: Boolean,
            value: false,
            notify: true
        },
        safeResponse: {
            type: Object,
            value: {},
            notify: true,
        },
        unsafeResponse: {
            type: Object,
            value: {},
            notify: true,
        },
        locationResponse: {
            type: Object,
            value: {},
            notify: true,
        },
        showMap: {
            type: Boolean,
            value: false,
            notify: true
        },
        showZikaMap: {
            type: Boolean,
            value: false,
            notify: true
        },
        latitude: {
            type: Number,
            value: 1.2956504,
            notify: true
        },
        longtitude: {
            type: Number,
            value: 103.8566111,
            notify: true
        },
        updateLocSuccess: {
          type: Boolean,
          value: false,
          notify: true
        }
    },
    listeners: {
		'google-map-ready': 'addCircles',
	},
	ready: function() {
        this.uuid = this.getUuid();
        console.log(this.uuid);
    },
    attached: function() {
        this.triggerLocationService();
    },
    getUuid: function() {
		url = window.location.href;
		name = ("uuid").replace(/[\[\]]/g, "\\$&");
		var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
			results = regex.exec(url);
		if (!results) return null;
		if (!results[2]) return '';
		return decodeURIComponent(results[2].replace(/\+/g, " ")); 
    },
    getParams: function(uuid) {
        return {uuid: uuid};
    },
    markAsSafe: function() {
        this.isSafe = true;
        this.marked = true;
        if (this.uuid && this.uuid !== "") {
            document.querySelector('#safe-ajax').generateRequest();
            console.log(this.safeResponse);
        }
    },
    markAsUnsafe: function() {
        this.isSafe = false;
        this.marked = true;
        if (this.uuid && this.uuid !== "") {
            document.querySelector('#unsafe-ajax').generateRequest();
            console.log(this.unsafeResponse);
        }
    },
    goToZikaView: function() {
        this.showZikaMap = true;
    },
    triggerLocationService: function(){
        if (navigator.geolocation) {
            navigator.geolocation.getCurrentPosition(function(position){
                this.latitude = position.coords.latitude;
                this.longtitude = position.coords.longitude;
console.log(position.coords);
                this.callForHelp();
            }.bind(this));
        }
    },
    callForHelp: function() {
        if(this.uuid && this.uuid !== "") {
            document.querySelector('#location-ajax').generateRequest();
        }
      this.updateLocSuccess = true;
    },
    getLocationParams: function(uuid){
        return {
            uuid: uuid,
            latitude: this.latitude,
            longitude: this.longtitude,
            timestamp: Date.now().toString()
        };
    },
	addCircles: function() {
		zikaMap = document.querySelector("#zika-map");
        var aljunied = {lat: 1.316566, lng: 103.88294};
        var aljCircle = new window.google.maps.Circle({
            strokeColor: '#FF0000',
            strokeOpacity: 0.8,
            strokeWeight: 2,
            fillColor: '#FF0000',
            fillOpacity: 0.35,
            map: zikaMap.map,
            center: aljunied,
            radius: 2500 
        });
        var bedok = {lat: 1.33733, lng: 103.93392};
        var bedokCircle = new window.google.maps.Circle({
            strokeColor: '#FF0000',
            strokeOpacity: 0.8,
            strokeWeight: 2,
            fillColor: '#FF0000',
            fillOpacity: 0.35,
            map: zikaMap.map,
            center: bedok,
            radius: 1000 
        });
        var yishun = {lat: 1.430087, lng: 103.83556};
        var yishCircle = new window.google.maps.Circle({
            strokeColor: '#FF0000',
            strokeOpacity: 0.8,
            strokeWeight: 2,
            fillColor: '#FF0000',
            fillOpacity: 0.35,
            map: zikaMap.map,
            center: yishun,
            radius: 1000 
        });
	},
});
