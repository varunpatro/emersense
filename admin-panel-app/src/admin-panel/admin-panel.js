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
      sheetUrl: {
        type: String,
        value: "https://docs.google.com/spreadsheets/d/1Q9xcqMUXLHF57-NvCQXSv2Q_NTT7L_rVsBqRNHL9E1c",
        notify: true
      }
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
    onEmergencyResponse: function(res){
      resp = res.detail.response;
      this.sheetUrl = resp.sheet_url;
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
