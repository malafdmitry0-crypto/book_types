import {tasks as original} from './catalog.mjs';
import {extra} from './catalog-extra.mjs';
import {exam} from './catalog-exam.mjs';
export const tasks=[...original,...extra,...exam];
