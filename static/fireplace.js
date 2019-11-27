$("#fireplace").submit(function(e){
    e.preventDefault();
    $.post("/fireplace", $("#fireplace").serialize())
      .done(function() {
        console.log("successfully set fireplace status");
      })
      .fail(function() {
        console.log("failed to set fireplace status");
      });
});
