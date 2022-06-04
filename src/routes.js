import { Router } from 'express';
import { infos } from "./countriesWithDial.service";


const router = Router();

router.get('/countries-with-dials', infos)

export { router }