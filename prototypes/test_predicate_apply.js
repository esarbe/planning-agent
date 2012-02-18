fs = require('fs');
logic = require('./predicate.js');
cp = require('./cp-problem.js');


function setup(data) {
  var problem = cp.buildProblem(data);

  return problem
} 

fs.readFile('/home/raphael/source/planning-agent/data/simple-world.json',
  'utf-8',
  function(err, data) {
    if (err) throw err;
    console.log(data)
    var data = JSON.parse(data);
    
    var problem = cp.buildProblem(data);
    var state = {}

    var in_hand = problem.predicates.in_hand;
    var on_table = problem.predicates.on_table;
    var not = logic.connectives.newNot([in_hand]);
    var and = logic.connectives.newAnd([not, on_table])
    
    in_hand.apply(state);
    on_table.apply(state);
    not.apply(state)
    
    console.log(state);
});


