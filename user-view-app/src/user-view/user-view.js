Polymer({
    is: 'user-view',
    properties: {
        userId: {
            type: String,
            value: "",
            notify: true
        },
        isSafe: {
            type: Boolean,
            value: false,
            notify: true
        },
    },
    markAsSafe: function() {
        this.isSafe = true;
        // Update server.
    },
    markAsUnsafe: function() {
        this.isSafe = false;
        // Update server.
    },
});
