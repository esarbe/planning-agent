
var Predicate = function (name, parameters) {
  parameters = typeof parameters !== 'undefined' ? parameters : [];

  this.name = name;
  this.parameters = parameters;
}

Predicate.prototype = Object.create(Object);
Predicate.prototype.constructor = Predicate
Predicate.prototype.apply = function (state, parameters, value) {
  value = typeof value !== 'undefined' ? value : true;
 
  if (value == true) {
    if (typeof state[this.name] === 'undefined'){
      state[this.name] = {}
    }
    state[this.name][parameters] = value
  } else {
    delete(state[this.name][parameters])   
  }
};

Predicate.prototype.isValidFor = function (state, params, value) {

}

Predicate.prototype.toString = function () {
  return ' (:' + this.name + ' ' + this.parameters + ') '
}

var Connective = function (name, children) {
  Predicate.call(this, name, {});
  this.children = children;
}

Connective.prototype = Object.create(Predicate);
Connective.prototype.constructor = Connective;
Connective.prototype.apply = function(world, value) {
  
  if (value == undefined) {
    value = true;
  }
  
  for (var predicate in this.children) {
    console.log('p:', typeof predicate)
    predicate.apply(world, this.children[predicate], this.eval(value))
  }
}

var newAnd = function (children) {

  var and = new Connective('and', children);
  and.eval = function (value) {
    return value;
  }
  return and;
}

var newNot = function (children) {
  var not = new Connective('not', children);
  not.eval = function (value) {
    return !value;
  }

  return not;
}

module.exports.connectives = {};
module.exports.connectives.newAnd = newAnd;
module.exports.connectives.newNot = newNot;
module.exports.Predicate = Predicate;
module.exports.Connective = Connective;


