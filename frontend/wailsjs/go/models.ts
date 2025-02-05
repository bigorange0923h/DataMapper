export namespace config {
	
	export class MappingsModel {
	    id: number;
	    sourceName: string;
	    sourceType: string;
	    targetName: string;
	    targetType: string;
	
	    static createFrom(source: any = {}) {
	        return new MappingsModel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.sourceName = source["sourceName"];
	        this.sourceType = source["sourceType"];
	        this.targetName = source["targetName"];
	        this.targetType = source["targetType"];
	    }
	}

}

