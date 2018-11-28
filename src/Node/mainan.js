var soap = require('soap');
var apiWSDL = 'http://brickset.com/api/v2.asmx?WSDL';

var key = 'pkWv-jYp1-ggRB';

//based on searching brickset.com
var minYear = 1950;

function getRandomInt (min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

function setKey(k) {
	key = k;
}

function getRandomSet() {

	//first, determine the year
	year = getRandomInt(minYear, (new Date()).getFullYear());
	console.log('chosen year', year);

	var p = new Promise(function(resolve, reject) {
		
		soap.createClient(apiWSDL, function(err, client) {
			if(err) throw new Error(err);

			var args = {
				apiKey:key,
				userHash:'',
				query:'',
				theme:'',
				subtheme:'',
				setNumber:'',
				year:year,
				owned:'',
				wanted:'',
				orderBy:'',
				pageSize:'2000',
				pageNumber:'1',
				userName:''
			}

			client.getSets(args, function(err, result) {
				if(err) reject(err);
				if(!result) {
					return getRandomSet();
                }

				var sets = result.getSetsResult.sets;
				console.log('i found '+sets.length+' results');
				if(sets.length) {
					var chosen = getRandomInt(0, sets.length-1);
					var set = sets[chosen];
					// now that we have a set, try to get more images
					if(set.additionalImageCount > 0) {
						client.getAdditionalImages({apiKey:key, setID:set.setID}, function(err, result) {
							if(err) reject(err);
							console.log('i got more images', result);
							set.additionalImages = result;
							resolve(set);
						});
					} else {
						resolve(set);
					}
				}
			});
		});
		
		
	});

	return p;

}

// exports.setKey = setKey;
// exports.getRandomSet = getRandomSet;

getRandomSet();