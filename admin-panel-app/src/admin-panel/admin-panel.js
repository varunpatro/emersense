Polymer({
    is: 'admin-panel',
    properties: {
      loggedIn: {
        type: Boolean,
        value: false,
        notify: true,
      },
      alertSent: {
        type: Boolean,
        value: false,
        notify: true,
      },
      userId: {
        type: String,
        value: "",
        notify: true,
      },
      response: {
        type: Object,
        value: {},
        notify: true
      },
    },
    doLogin: function() {
        console.log(this.$);
        var email = document.querySelector("#email-input").value;
        var password = document.querySelector("#password-input").value;
        // TODO Send this to the server.
        console.log(email);
        console.log(password);
        this.loggedIn = true;
    },
    sendAlert: function() {
        if (!this.alertSent) {
            var sender = document.querySelector("#send-alert-ajax");
            sender.generateRequest();       
            this.alertSent = true;
        }
    },
    computeTitle: function(userId) {
        console.log('Here');
        if (!userId || userId === "") {
            return "EmerSense";
        } else {
            return "EmerSense for " + userId;
        }
    }
});
