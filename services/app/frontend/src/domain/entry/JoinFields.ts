import {Field} from "../category/Category";

export default function JoinFields(bib: Field[], cite: Field[]): Field[] {
    const bibNames = bib.map(f => f.Name);
    cite.forEach(f => {
        if (bibNames.indexOf(f.Name) === -1) {
            bib.push(f);
        }
    });

    return bib;
}
