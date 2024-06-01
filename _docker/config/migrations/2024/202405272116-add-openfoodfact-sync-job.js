const KEY = 'OFF_SYNC_JOB';
const execute = async (db, context = {}) => {
  const { uuid } = context;
  const job = {
    _id: uuid(),
    cronExpression: "0 0 3 * * *",
    createdAt: new Date(),
    key: KEY,
    name: 'OpenFoodFacts: Daily Sync',
    description: 'refresh the OFF database every day',
    running: false,
    disabled: false,
    params: {
      endpoint:
        'https://static.openfoodfacts.org/data/en.openfoodfacts.org.products.csv.gz',
      separator: "\t",
      gzip: true,
      parallelism: 4,
      batchSize100Ms: 20, // 4*20*10 = 800req/sec
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
