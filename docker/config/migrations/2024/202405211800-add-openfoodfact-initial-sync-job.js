const KEY = 'OFF_INITIAL_SYNC_JOB';
const execute = async (db, context = {}) => {
  const { now, uuid } = context;
  const job = {
    _id: uuid(),
    // cronExpression: "0 0 0 * * *",
    createdAt: now(),
    specificDate: now(),
    key: KEY,
    name: 'OpenFoodFacts: Initial Sync',
    description: 'fill the FF database with OF data, using gzipped CSV',
    params: {
      endpoint:
        'https://static.openfoodfacts.org/data/en.openfoodfacts.org.products.csv.gz',
      separator: '\t',
    },
  };

  const jobCollection = await db.collection('job');
  await jobCollection.insertOne(job);
};

const rollback = async (db, _context = {}) => {
  const collection = await db.collection('job');
  await collection.deletOne({ key: KEY });
};

module.exports = {
  targetDatabases: ['factsfood'],
  description: 'Open Food Facts Initial Sync job',
  rollback,
  execute,
};
