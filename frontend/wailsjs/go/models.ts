export namespace models {
	
	export class CatalogItem {
	    id: number;
	    name: string;
	    stars: number;
	    owner: string;
	    git_url: string;
	    avatar_url: string;
	    is_favourite: boolean;
	
	    static createFrom(source: any = {}) {
	        return new CatalogItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.stars = source["stars"];
	        this.owner = source["owner"];
	        this.git_url = source["git_url"];
	        this.avatar_url = source["avatar_url"];
	        this.is_favourite = source["is_favourite"];
	    }
	}
	export class FavouriteItem {
	    id: number;
	    catalog_item: CatalogItem;
	    current_release: string;
	    latest_release: string;
	
	    static createFrom(source: any = {}) {
	        return new FavouriteItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.catalog_item = this.convertValues(source["catalog_item"], CatalogItem);
	        this.current_release = source["current_release"];
	        this.latest_release = source["latest_release"];
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
	export class TagResponse {
	    current: string;
	    latest: string;
	
	    static createFrom(source: any = {}) {
	        return new TagResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.current = source["current"];
	        this.latest = source["latest"];
	    }
	}

}

