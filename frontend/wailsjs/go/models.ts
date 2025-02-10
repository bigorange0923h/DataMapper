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
	export class MenuItem {
	    key: string;
	    title: string;
	    children?: MenuItem[];
	
	    static createFrom(source: any = {}) {
	        return new MenuItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.title = source["title"];
	        this.children = this.convertValues(source["children"], MenuItem);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

