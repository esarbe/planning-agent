fs = require('fs');
logic = require('./predicate.js');
cp = require('./cp-problem.js');

repl = require("repl")

function setup(data) {
  var problem = cp.buildProblem(data);

  return problem
}

apply_action = function(state, action, params) {

  if (action.precondition.isValidFor(state, params)){
    action.effect.apply(state, params);
  } else {
    console.log("Action '", action.name, "' not valid for given state and parameters");  
  }
}

fs.readFile('/home/raphael/source/planning-agent/data/simple-world.json',
  'utf-8',
  function(err, data) {
    if (err) throw err;
    console.log(data)
    var data = JSON.parse(data);
    
    var problem = cp.buildProblem(data);
    var state = {}


    var isDay = new logic.Predicate( 'isDay' ); 
    var exists = new logic.Predicate( 'exists', ['obj']);
    var not_exists = logic.connectives.newNot([exists]);

    var action = {
      name : 'create_object',
      params : ['obj'],
      effect : exists,
      precondition : not_exists
    }

    apply_action(state, action, ['apple']);


    console.log(JSON.stringify(state));
    
});


