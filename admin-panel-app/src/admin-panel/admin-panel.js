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
            // TODO Do the ajax call.
            this.alertSent = true;
        }
    },
});
